package cdk

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/netvolart/joven/internal/iac"
)

type CDKPackage iac.Package

type PackageJson struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func parsePackageJson(data []byte) []CDKPackage {
	var packages []CDKPackage
	// open package.json file and unmarshal to struct
	file := PackageJson{}
	err := json.Unmarshal(data, &file)
	if err != nil {
		log.Fatalf("error unmarshaling package JSON: %v", err)
	}
	for key := range file.DevDependencies {
		packages = append(packages, CDKPackage{Name: key, LocalVersion: file.DevDependencies[key]})
	}
	for key := range file.Dependencies {
		packages = append(packages, CDKPackage{Name: key, LocalVersion: file.Dependencies[key]})
	}
	return packages
}

type TreeNode struct {
	ID            string              `json:"id"`
	Path          string              `json:"path"`
	ConstructInfo ConstructInfo       `json:"constructInfo"`
	Children      map[string]TreeNode `json:"children,omitempty"`
}

type Tree struct {
	Version string   `json:"version"`
	Tree    TreeNode `json:"tree"`
}

type ConstructInfo struct {
	Fqn     string `json:"fqn"`
	Version string `json:"version"`
}

type Children struct {
	ID            string        `json:"id"`
	Path          string        `json:"path"`
	ConstructInfo ConstructInfo `json:"constructInfo"`
}

func getNodeCDKConstructs(data []byte) []ConstructInfo {
	tree := Tree{}
	err := json.Unmarshal(data, &tree)
	if err != nil {
		log.Fatalf("error unmarshaling tree JSON: %v", err)
	}
	constructs := []ConstructInfo{}

	for _, value := range tree.Tree.Children {
		children := Children{}
		childJson, err := json.Marshal(value)

		if err != nil {
			log.Fatalf("error marshaling tree JSON: %v", err)
		}
		err = json.Unmarshal(childJson, &children)
		if err != nil {
			log.Fatalf("error unmarshaling tree JSON: %v", err)
		}
		constructs = append(constructs, children.ConstructInfo)

	}
	return constructs

}

func formNodeCDKPackages(constructs []ConstructInfo) []CDKPackage {
	packages := []CDKPackage{}
	for _, construct := range constructs {
		packages = append(packages, CDKPackage{
			Name:         clearFqn(construct.Fqn),
			LocalVersion: construct.Version,
		})
	}
	return packages

}

func clearFqn(Fqn string) string {
	lib := strings.Split(Fqn, ".")[0]
	if lib != "" {
		return lib
	}
	return ""
}

func removeDuplicates(packages []CDKPackage) []CDKPackage {
	keys := make(map[string]bool)
	list := []CDKPackage{}
	for _, entry := range packages {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}
	return list
}


func (p *CDKPackage) getPackageNpmInfo() {
	cmd := exec.Command("npm", "view", p.Name, "--json")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %s", err)

	}

	type npmView struct {
		Name     string `json:"name"`
		DistTags struct {
			Latest string `json:"latest"`
		} `json:"dist-tags"`
		Dist struct {
			Tarball string `json:"tarball"`
		} `json:"dist"`
		Version string `json:"version"`
	}
	view := npmView{}
	err = json.Unmarshal(output, &view)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)

	}
	// Check if valid semver
	_, err = semver.NewVersion(view.Version)
	if err != nil {
		log.Fatalf("Version is not in Semver format: %s", err)
	}

	p.LatestVersion = view.DistTags.Latest
	p.Link = view.Dist.Tarball

}

func (p *CDKPackage) setNpmPackageType() {
	if strings.Contains(p.Link, "gitlab.com") {
		p.Type = "gitlab"
	} else if strings.Contains(p.Link, "registry.npmjs.org") {
		p.Type = "npmjs"
	} else {
		p.Type = "unknown"
	}
}

func (p *CDKPackage) setOutdated() {
	latestVersion, err := semver.NewVersion(p.LatestVersion)
	if err != nil {
		log.Fatalf("Error parsing version: %s", err)
	}
	localVersion, err := semver.NewVersion(p.LocalVersion)
	if err != nil {
		log.Fatalf("Error parsing version: %s", err)
	}
	if latestVersion.GreaterThan(localVersion) {
		p.Outdated = true
	}
}
