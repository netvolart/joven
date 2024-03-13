package terraform

import (
	"fmt"
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
	localModules, err := GetLocalModules()
	if err != nil {
		return nil, err
	}
	var resultModules []*TerraformModule
	for _, localModule := range localModules.Modules {
		if localModule.Source == "" {
			continue
		}
		if !strings.Contains(localModule.Source, "gitlab") {
			continue
		}
		namesList := strings.Split(localModule.Source, "/")

		nameAndVendor := fmt.Sprintf("%s/%s", namesList[len(namesList)-2], namesList[len(namesList)-1])
		url, err := createModuleGitlabUrl(c, nameAndVendor)
		if err != nil {
			return nil, err
		}

		gitlabModulesVersions, err := getModuleVersionsFromGitLab(c, url)
		if err != nil {
			return nil, err
		}
		result, err := clearOldVersions(gitlabModulesVersions)

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
