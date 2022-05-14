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
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	fmt.Fprintf(os.Stdout, "%d items reported\n", len(results))
}
