package analyzer

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func ReportStdout(results []AnaResult) {
	data := [][]string{}
	for _, r := range results {
		price := r.Work.Price
		if r.Work.Discount != 0 {
			price = r.Work.Discount
		}
		row := []string{
			r.Work.Name,
			r.Work.URL,
			fmt.Sprintf("%d yen", price),
			r.Report,
		}
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "URL", "Yen", "Report"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
