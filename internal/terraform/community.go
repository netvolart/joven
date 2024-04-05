package terraform

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/netvolart/joven/internal/iac"
)

func getModuleVersionsFromRegistry(url string, localModule LocalModule) (module iac.Package, Error error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return module, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return module, err
	}
	defer resp.Body.Close()

	type Response struct {
		ID      string `json:"id"`
		Version string `json:"version"`
		Link    string `json:"source"`
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return module, err
	}

	module.Name = localModule.Source
	module.LocalVersion = localModule.Version
	module.LatestVersion = response.Version
	module.Link = response.Link

	return module, nil

}

func createModuleCommunityUrl(moduleName string) (string, error) {
	if moduleName == "" {
		return "", ErrorPageNumberEmpty
	}
	baseURL, err := url.Parse("https://registry.terraform.io/v1/modules/")
	if err != nil {
		return "", err
	}
	parts := strings.Split(moduleName, "/")
	moduleWithoutRegistry := strings.Join(parts[1:], "/")
	pathURL, err := url.Parse(moduleWithoutRegistry)
	if err != nil {
		return "", err
	}
	return baseURL.ResolveReference(pathURL).String(), nil

}
