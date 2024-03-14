package terraform

import (
	"encoding/json"
	"os"
)

type LocalModule struct {
	Source  string `json:"Source"`
	Version string `json:"Version"`
}

type LocalModules struct {
	Modules []LocalModule `json:"Modules"`
}

func GetLocalModules() (*LocalModules, error) {
	configData, err := os.ReadFile(".terraform/modules/modules.json")
	if err != nil {
		return nil, err
	}
	var localModules LocalModules
	err = json.Unmarshal(configData, &localModules)
	if err != nil {
		return nil, err
	}
	var cleanLocalModules LocalModules
	for _, module := range localModules.Modules {
		if module.Source != "" {
			cleanLocalModules.Modules = append(cleanLocalModules.Modules, module)
		}
	}
	return &localModules, nil
}
