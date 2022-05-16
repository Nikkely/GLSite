package analyzer

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Reporter interface {
	Report(methodName string, results []AnaResult)
}

type reportStdoutTable struct{}
type reportStdoutSimple struct{}

func NewReportStdoutTable() reportStdoutTable {
	return reportStdoutTable{}
}

func NewReportStdoutSimple() reportStdoutSimple {
	return reportStdoutSimple{}
}

func (r reportStdoutTable) Report(methodName string, results []AnaResult) {
	fmt.Fprintln(os.Stdout, methodName)
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

func (r reportStdoutSimple) Report(methodName string, results []AnaResult) {
	for i, a := range results {
		fmt.Fprintf(os.Stdout, "\n%s_%d\n%s %s\n%s\n%s\n", methodName, i+1, a.Work.Maker.Name, a.Work.Name, a.Work.URL, a.Report)
	}
}
