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
			fmt.Println(nameAndVendor)
			fmt.Println(gitlabModule.Name)

			if (nameAndVendor == gitlabModule.Name) && strings.Contains(localModule.Source, "gitlab") {
				fmt.Println("here")
				mod := NewTerraformModule(gitlabModule.Name, localModule.Version, gitlabModule.LatestVersion)
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
		if  latestVersion.GreaterThan(localVersion) {
			outdatedModules = append(outdatedModules, module)
		}
	}
	return outdatedModules, nil
}