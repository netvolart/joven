package terraform

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Masterminds/semver"
	config "github.com/volkovartem/joven/config"
)

type TerraformModule struct {
	Name          string
	LocalVersion  string
	LatestVersion string
}

type Response struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	// Add other fields as needed based on the JSON response
}

func NewTerraformModule(name string, version string, latestVersion string) *TerraformModule {
	return &TerraformModule{
		Name:          name,
		LocalVersion:  version,
		LatestVersion: latestVersion}
}

func createGitLabUrl(c *config.Config, page string) (string, error) {
	baseURL, err := url.Parse("https://gitlab.com/api/v4/groups/")
	if err != nil {
		return "", err
	}
	pathURL, err := url.Parse(fmt.Sprintf("%s/packages?package_type=terraform_module&pagination=keyset&page=%s&per_page=100&sort=asc", c.Groups[0], page))
	if err != nil {
		return "", err
	}
	return baseURL.ResolveReference(pathURL).String(), nil
}

func makeGiLabModulesRequest(c *config.Config, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Add headers to the request
	req.Header.Add("PRIVATE-TOKEN", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func gitlabModulesRequest(c *config.Config) (*[]Response, error) {
	url, err := createGitLabUrl(c, "1")
	if err != nil {
		return nil, err
	}
	resp, err := makeGiLabModulesRequest(c, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var fullResponses []Response
	var responses []Response
	err = json.NewDecoder(resp.Body).Decode(&responses)

	fullResponses = append(fullResponses, responses...)

	totalPages, err := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Total pages: %v\n", totalPages)
	for i := 2; i <= totalPages; i++ {
		url, err := createGitLabUrl(c, strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		resp, err := makeGiLabModulesRequest(c, url)
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(resp.Body).Decode(&responses)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		fullResponses = append(fullResponses, responses...)
	}

	return &fullResponses, nil

}

func GetModulesFromGitlab(c *config.Config) ([]*TerraformModule, error) {
	log.Printf("Getting modules from GitLab")
	responses, err := gitlabModulesRequest(c)
	if err != nil {
		log.Printf("Error getting modules from GitLab: %v", err)

	}
	var modules []*TerraformModule
	for _, response := range *responses {
		module := NewTerraformModule(response.Name, "", response.Version)
		modules = append(modules, module)
	}
	clearOldVersions(modules)
	return modules, nil
}

func clearOldVersions(modules []*TerraformModule) ([]*TerraformModule, error) {
	latestModules := make(map[string]*TerraformModule)
	for _, module := range modules {
		if module == nil {
			continue
		}
		if existingModule, ok := latestModules[module.Name]; ok {

			moduleVersion, err := semver.NewVersion(module.LatestVersion)
			if err != nil {
				return nil, err
			}
			mapVersion, err := semver.NewVersion(existingModule.LatestVersion)
			if err != nil {
				return nil, err
			}
			if moduleVersion.GreaterThan(mapVersion) {
				latestModules[module.Name].LatestVersion = module.LatestVersion
			}
		} else {
			latestModules[module.Name] = module
		}
	}

	var latestModulesSlice []*TerraformModule
	for _, module := range latestModules {
		latestModulesSlice = append(latestModulesSlice, module)
	}

	return latestModulesSlice, nil
}
