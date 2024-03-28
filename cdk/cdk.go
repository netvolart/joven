package cdk

import (
	"fmt"
	"log"
	"os"
)

func CompareCDKConstructs() {
	lang, err := detectLanguage()
	if err != nil {
		log.Fatalf("error detecting language: %v", err)
	}
	switch lang {
	case "typescript":
		localCdkTree, err := os.ReadFile("cdk.out/tree.json")
		if err != nil {
			log.Fatalf("error reading tree JSON: %v", err)
		}
		constructs := getNodeCDKConstructs(localCdkTree)
		packages := formNodeCDKPackages(constructs)
		packagesWithoutDuplicates := removeDuplicates(packages)
		for k, pack := range packagesWithoutDuplicates {
			pack.getPackageNpmInfo()
			pack.setNpmPackageType()
			pack.setOutdated()
			packagesWithoutDuplicates[k] = pack
		}
		fmt.Println(packagesWithoutDuplicates)
	}
}