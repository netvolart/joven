package terraform

import (
	"encoding/json"
	"fmt"
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
