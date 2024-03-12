package terraform

import (
	"fmt"
	"strings"
)

func MergeModules(gitlabModules []*TerraformModule, localModules LocalModules) []*TerraformModule {
	var resultModules []*TerraformModule
	for _, localModule := range localModules.Modules {
		if localModule.Source == "" {
			continue
		}
		namesList := strings.Split(localModule.Source, "/")
		//	fmt.Println(namesList)
		// fmt.Println(namesList[len(namesList)-1])

		nameAndVendor := fmt.Sprintf("%s/%s", namesList[len(namesList)-1], namesList[len(namesList)-2])
		for _, gitlabModule := range gitlabModules {

			if (nameAndVendor == gitlabModule.Name) && strings.Contains(localModule.Source, "gitlab") {
				fmt.Println("here")
				mod := NewTerraformModule(gitlabModule.Name, localModule.Version, gitlabModule.LatestVersion)
				resultModules = append(resultModules, mod)
			}
		}
	}
	return resultModules
}
