package terraform

import (
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/volkovartem/joven/config"
)

func findOutdated(modules []*TerraformModule) ([]*TerraformModule, error) {
	markedModules := []*TerraformModule{}

	for _, module := range modules {
		latestVersion, err := semver.NewVersion(module.LatestVersion)
		if err != nil {
			return nil, err
		}
		localVersion, err := semver.NewVersion(module.LocalVersion)
		if err != nil {
			return nil, err
		}
		if latestVersion.GreaterThan(localVersion) {
			module.Outdated = true
			markedModules = append(markedModules, module)
			continue
		}
		markedModules = append(markedModules, module)
	}
	return markedModules, nil
}

func CompareGitLabModules(c *config.Config) ([]*TerraformModule, error) {
	configData, err := os.ReadFile(".terraform/modules/modules.json")

	if err != nil {
		return nil, err
	}
	localModulesResult, err := GetLocalModules(configData)
	if err != nil {
		return nil, err
	}
	localModules := setModulesSourceType(localModulesResult)
	var resultModules []*TerraformModule
	for _, localModule := range localModules.Modules {
		var remoteModules []*TerraformModule
		if (localModule.Type == "gitlab") {
			namesList := strings.Split(localModule.Source, "/")

			nameAndVendor := fmt.Sprintf("%s/%s", namesList[len(namesList)-2], namesList[len(namesList)-1])
			url, err := createModuleGitlabUrl(c, nameAndVendor)
			if err != nil {
				return nil, err
			}
	
			remoteModules, err = getModuleVersionsFromGitLab(c, url)
			if err != nil {
				return nil, err
			}
		} else if (localModule.Type == "community") {
			continue
		}

		result, err := clearOldVersions(remoteModules)

		if err != nil {
			return nil, err
		}

		mod := NewTerraformModule(result[0].Name, localModule.Version, result[0].LatestVersion, result[0].Link, false)
		resultModules = append(resultModules, mod)
	}

	withMarkedOutdated, err := findOutdated(resultModules)
	if err != nil {
		return nil, err
	}
	return withMarkedOutdated, nil
}
