package terraform

import (
	"encoding/json"
	"strings"
)

type LocalModule struct {
	Key     string `json:"Key"`
	Source  string `json:"Source"`
	Version string `json:"Version"`
	Type    string
	Nested  bool
}

func (m *LocalModule) checkIfNested() {
	if strings.Contains(m.Source, "//") || strings.Contains(m.Key, ".") {
		m.Nested = true
	}
}

func setModulesSourceType(modules *LocalModules) *LocalModules {
	var modulesWithTypes LocalModules
	for _, m := range modules.Modules {
		var moduleType string
		if strings.Contains(m.Source, "gitlab.com") {
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
			Nested:  m.Nested,
		}
		modulesWithTypes.Modules = append(modulesWithTypes.Modules, mod)

	}
	return &modulesWithTypes

}

type LocalModules struct {
	Modules []LocalModule `json:"Modules"`
}

func getLocalModules(data []byte) (*LocalModules, error) {
	var localModules LocalModules
	err := json.Unmarshal(data, &localModules)
	if err != nil {
		return nil, err
	}
	var cleanLocalModules LocalModules
	for _, module := range localModules.Modules {
		// Avoid empty module sources
		if module.Source == "" {
			continue
		}
		module.checkIfNested()
		// Nested modules (modules in modules) are not supported
		// To keep output readable, we skip them
		if module.Nested {
			continue
		}
		if strings.Contains(module.Source, "//") {
			parentModuleSlice := strings.Split(module.Source, "//")
			module.Source = parentModuleSlice[len(parentModuleSlice)-1]
		}
		cleanLocalModules.Modules = append(cleanLocalModules.Modules, module)
	}
	return &cleanLocalModules, nil
}
