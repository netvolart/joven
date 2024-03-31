package terraform

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/netvolart/joven/internal/config"
	"github.com/netvolart/joven/internal/iac"
)

func findOutdated(modules []*iac.Package) ([]*iac.Package, error) {
	markedModules := []*iac.Package{}

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

func CompareGitLabModules(c *config.Config, localModulesData []byte) ([]*iac.Package, error) {
	// parse local modules
	localModulesResult, err := getLocalModules(localModulesData)
	if err != nil {
		return nil, err
	}
	// set types for local modules (gitlab or community)
	localModules := setModulesSourceType(localModulesResult)
	var resultModules []*iac.Package

	for _, localModule := range localModules.Modules {
		var remoteModules []*iac.Package
		if localModule.Type == "gitlab" {
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
			result, err := clearOldVersions(remoteModules)

			if err != nil {
				return nil, err
			}
			mod := iac.NewPackage(result[0].Name, localModule.Version, result[0].LatestVersion, result[0].Link, false)
			resultModules = append(resultModules, mod)

		} else if localModule.Type == "community" {

			url, err := createModuleCommunityUrl(localModule.Source)
			if err != nil {
				return nil, err
			}
			communityModule, err := getModuleVersionsFromRegistry(url, localModule)
			if err != nil {
				return nil, err
			}
			resultModules = append(resultModules, &communityModule)
		}

	}

	withMarkedOutdated, err := findOutdated(resultModules)
	if err != nil {
		return nil, err
	}
	return withMarkedOutdated, nil
}
