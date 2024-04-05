package terraform

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/netvolart/joven/internal/helper"
	"github.com/netvolart/joven/internal/iac"
)

func Print(w io.Writer, modules []*iac.Package) {
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(w)

	t.AppendHeader(table.Row{"Module", "Local Version", "Latest Version", "Status"})
	for _, module := range modules {

		r := table.Row{
			helper.AddColor(module.Outdated, module.Name),
			helper.AddColor(module.Outdated, module.LocalVersion),
			helper.AddColor(module.Outdated, module.LatestVersion),
			helper.AddColor(module.Outdated, helper.HumanReadableOutdated(module.Outdated)),
		}
		link := table.Row{fmt.Sprintf("Latest version -> %s", helper.AddColor(module.Outdated, module.Link))}

		t.AppendRow(r)
		t.AppendRow(link)
		t.AppendSeparator()
	}
	t.Render()

}
