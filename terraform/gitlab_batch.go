package terraform

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	config "github.com/netvolart/joven/config"
)

type Response struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Links   struct {
		WebPath string `json:"web_path"`
	} `json:"_links"`
}

func createGitLabUrl(c *config.Config, page string) (string, error) {
	if page == "" {
		return "", ErrorPageNumberEmpty
	}
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

func makeGiLabModulesRequest(c *config.Config, url string) (modulesResp *[]Response, totalCount int, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	// Add headers to the request
	req.Header.Add("PRIVATE-TOKEN", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	var responses []Response
	err = json.NewDecoder(resp.Body).Decode(&responses)
	if err != nil {
		return nil, 0, err
	}
	totalPages, err := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
	if err != nil {
		return nil, 0, err
	}
	return &responses, totalPages, nil
}

func downloadModulesMetadata(c *config.Config) ([]Response, error) {
	url, err := createGitLabUrl(c, "1")
	if err != nil {
		return nil, err
	}
	responses, totalPages, err := makeGiLabModulesRequest(c, url)
	if err != nil {
		return nil, err
	}

	var fullResponses []Response

	fullResponses = append(fullResponses, *responses...)
	resultChannel := make(chan *[]Response)
	for i := 2; i <= totalPages; i++ {

		url, err := createGitLabUrl(c, strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		go func(c *config.Config, url string) {
			responses, _, err := makeGiLabModulesRequest(c, url)
			if err != nil {
				log.Println(err)
				return
			}
			resultChannel <- responses
		}(c, url)

		fullResponses = append(fullResponses, *responses...)
	}
	for i := 2; i <= totalPages; i++ {
		r := <-resultChannel
		fullResponses = append(fullResponses, *r...)
	}
	return fullResponses, nil
}

func getModulesFromGitlab(c *config.Config) ([]*TerraformModule, error) {
	responses, err := downloadModulesMetadata(c)
	if err != nil {
		log.Printf("Error getting modules from GitLab: %v", err)

	}
	var modules []*TerraformModule
	for _, response := range responses {
		link := "https://gitlab.com" + response.Links.WebPath
		module := NewTerraformModule(response.Name, "", response.Version, link, false)
		modules = append(modules, module)
	}

	cleared, err := clearOldVersions(modules)
	if err != nil {
		log.Printf("Unable to clean modules: %v", err)
	}

	return cleared, nil
}
