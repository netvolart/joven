package terraform

import (
	"fmt"
	"io"

	"github.com/alexeyco/simpletable"
)

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

type TerraformModule struct {
	Name          string
	LocalVersion  string
	LatestVersion string
	Link          string
	Outdated      bool
	Type          string
}



func NewTerraformModule(name, localVersion, latestVersion, link string, outdated bool) *TerraformModule {
	return &TerraformModule{
		Name:          name,
		LocalVersion:  localVersion,
		LatestVersion: latestVersion,
		Link:          link,
		Outdated:      outdated,
	}
}

type TerraformModules []*TerraformModule


func (modules TerraformModules) Print(w io.Writer) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Module"},
			{Align: simpletable.AlignCenter, Text: "Local Version"},
			{Align: simpletable.AlignCenter, Text: "Latest Version"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Link"},
		},
	}

	for _, module := range modules {

		r := []*simpletable.Cell{

			{Text: addColor(module.Outdated, module.Name)},
			{Text: addColor(module.Outdated, module.LocalVersion)},
			{Text: addColor(module.Outdated, module.LatestVersion)},
			{Text: addColor(module.Outdated, humanReadableOutdated(module.Outdated))},
			{Text: addColor(module.Outdated, module.Link)},
		}

		table.Body.Cells = append(table.Body.Cells, r)

	}
	table.Println()

}

func humanReadableOutdated(outdated bool) string {
	if outdated {
		return "Outdated"
	}
	return ""

}

func addColor(outdated bool, s string) string {
	if outdated {
		return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
	}
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)

}
