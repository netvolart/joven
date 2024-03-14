package terraform

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Masterminds/semver"
	"github.com/volkovartem/joven/config"
)

var ErrorPageNumberEmpty = errors.New("Page can't be empty")

func createModuleGitlabUrl(c *config.Config, moduleName string) (string, error) {
	if moduleName == "" {
		return "", ErrorPageNumberEmpty
	}
	baseURL, err := url.Parse("https://gitlab.com/api/v4/groups/")
	if err != nil {
		return "", err
	}
	pathURL, err := url.Parse(fmt.Sprintf("%s/packages?package_type=terraform_module&package_name=%s&sort=asc", c.Groups[0], moduleName))
	if err != nil {
		return "", err
	}
	return baseURL.ResolveReference(pathURL).String(), nil

}

func getModuleVersionsFromGitLab(c *config.Config, url string) (modules []*TerraformModule, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("PRIVATE-TOKEN", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type Response struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Links   struct {
			WebPath string `json:"web_path"`
		} `json:"_links"`
	}
	var responses []Response
	err = json.NewDecoder(resp.Body).Decode(&responses)
	if err != nil {
		return nil, err
	}

	for _, response := range responses {
		link := "https://gitlab.com" + response.Links.WebPath
		module := NewTerraformModule(response.Name, "", response.Version, link, false)
		modules = append(modules, module)
	}

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
