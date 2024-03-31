package iac

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
)

type Package struct {
	Name          string
	LocalVersion  string
	LatestVersion string
	Link          string
	Outdated      bool
	Type          string
}

func NewPackage(name, localVersion, latestVersion, link string, outdated bool) *Package {
	return &Package{
		Name:          name,
		LocalVersion:  localVersion,
		LatestVersion: latestVersion,
		Link:          link,
		Outdated:      outdated,
	}
}

func Print(w io.Writer, modules []*Package) {
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(w)

	t.AppendHeader(table.Row{"Module", "Local Version", "Latest Version", "Status"})
	for _, module := range modules {

		r := table.Row{
			addColor(module.Outdated, module.Name),
			addColor(module.Outdated, module.LocalVersion),
			addColor(module.Outdated, module.LatestVersion),
			addColor(module.Outdated, humanReadableOutdated(module.Outdated)),
		}
		link := table.Row{fmt.Sprintf("Latest version -> %s", addColor(module.Outdated, module.Link))}

		t.AppendRow(r)
		t.AppendRow(link)
		t.AppendSeparator()
	}
	t.Render()

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
