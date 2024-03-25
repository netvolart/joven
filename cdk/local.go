package cdk

import (
	"fmt"
	"os"
	"path/filepath"
)

func detectLanguage() (lang string, err error) {
	root := "."
	patterns := map[string]string{
		"*.csproj":     "dotnet",
		"package.json": "typescript",
	}

	fileNames := []string{}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileNames = append(fileNames, path)

		return nil

	})

	for _, fileName := range fileNames {
		for pattern := range patterns {
			matched, err := filepath.Match(pattern, filepath.Base(fileName))
			if err != nil {
				return "", err
			}
			if matched {
				fmt.Println(patterns[pattern])
				return patterns[pattern], nil
			}
		}
	}

	// if matched {
	// 	fmt.Println(patterns[pattern])
	// 	lang = patterns[pattern]

	// }
	// for _, pattern := range patterns {

	// 	if err != nil {
	// 		fmt.Printf("walk error [%v]\n", err)
	// 	}
	// }
	// return lang, nil
	return "", nil
}
