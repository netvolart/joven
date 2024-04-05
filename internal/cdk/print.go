package cdk

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/netvolart/joven/internal/helper"
)

func Print(w io.Writer, packages []CDKPackage) {
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(w)

	t.AppendHeader(table.Row{"Construct", "Local Version", "Latest Version", "Status"})
	for _, pack := range packages {

		r := table.Row{
			helper.AddColor(pack.Outdated, pack.Name),
			helper.AddColor(pack.Outdated, pack.LocalVersion),
			helper.AddColor(pack.Outdated, pack.LatestVersion),
			helper.AddColor(pack.Outdated, helper.HumanReadableOutdated(pack.Outdated)),
		}
		t.AppendRow(r)
		t.AppendSeparator()
	}
	t.Render()

}
