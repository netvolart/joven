package terraform

import (
	"encoding/json"
	"strings"
)

type LocalModule struct {
	Source  string `json:"Source"`
	Version string `json:"Version"`
	Type    string
}



func setModulesSourceType(modules *LocalModules) *LocalModules {
	var modulesWithTypes LocalModules
	for _, m := range modules.Modules {
		var moduleType string
		if strings.Contains(m.Source, "gitlab") {
			moduleType = "gitlab"
		} else if strings.Contains(m.Source, "registry.terraform.io") {
			moduleType = "community"
		} else {
			moduleType = "unknown"
		}
		mod := LocalModule{
			Source:  m.Source,
			Version: m.Version,
			Type:    moduleType,
		}
		modulesWithTypes.Modules = append(modulesWithTypes.Modules, mod)

	}
	return &modulesWithTypes

}

type LocalModules struct {
	Modules []LocalModule `json:"Modules"`
}

func GetLocalModules(data []byte) (*LocalModules, error) {
	var localModules LocalModules
	err := json.Unmarshal(data, &localModules)
	if err != nil {
		return nil, err
	}
	var cleanLocalModules LocalModules
	for _, module := range localModules.Modules {
		if module.Source == "" {
			continue
		}
		cleanLocalModules.Modules = append(cleanLocalModules.Modules, module)
	}
	return &cleanLocalModules, nil
}
