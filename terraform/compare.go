package terraform

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
)

func MergeModules(gitlabModules []*TerraformModule, localModules LocalModules) []*TerraformModule {
	var resultModules []*TerraformModule
	for _, localModule := range localModules.Modules {
		if localModule.Source == "" {
			continue
		}
		namesList := strings.Split(localModule.Source, "/")

		nameAndVendor := fmt.Sprintf("%s/%s", namesList[len(namesList)-2], namesList[len(namesList)-1])
		for _, gitlabModule := range gitlabModules {
			if (nameAndVendor == gitlabModule.Name) && strings.Contains(localModule.Source, "gitlab") {
				mod := NewTerraformModule(gitlabModule.Name, localModule.Version, gitlabModule.LatestVersion, gitlabModule.Link, false)
				resultModules = append(resultModules, mod)
			}
		}
	}
	return resultModules
}

func returnOutdated(modules []*TerraformModule) ([]*TerraformModule, error) {
	outdatedModules := []*TerraformModule{}

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
			outdatedModules = append(outdatedModules, module)
		}
	}
	return outdatedModules, nil
}

func FindOutdated(modules []*TerraformModule) ([]*TerraformModule, error) {
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