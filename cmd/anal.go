package cmd

import (
	"log"

	"github.com/Nikkely/GLSite/analyzer"
	"github.com/spf13/cobra"
)

const (
	analyzeDirPathFlg     = "dir-path"
	analyzeDLThresholdFlg = "dl-th"
	analyzeOutputFormat   = "format"
)

func init() {
	rootCmd.AddCommand(AnalCmd)

	AnalCmd.PersistentFlags().StringP(analyzeDirPathFlg, "d", "output", "path to output direcotry")
	AnalCmd.PersistentFlags().Int(analyzeDLThresholdFlg, 50, "path to output direcotry")
	AnalCmd.PersistentFlags().StringP(analyzeOutputFormat, "f", "table", "output format")
}

var AnalCmd = &cobra.Command{
	Use:   "anal",
	Short: "Analyzer",
	Long:  `Analyze datas`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString(analyzeDirPathFlg)
		if err != nil {
			log.Fatalln(err.Error())
		}
		dlth, err := cmd.Flags().GetInt(analyzeDLThresholdFlg)
		if err != nil {
			log.Fatalln(err.Error())
		}
		format, err := cmd.Flags().GetString(analyzeOutputFormat)
		if err != nil {
			log.Fatalln(err.Error())
		}

		methods := []analyzer.AnalyzeMethod{
			analyzer.NewChangePrice(),
			analyzer.NewisNew(),
			analyzer.NewChangeDL(dlth),
		}

		var reporter analyzer.Reporter
		switch format {
		case "table":
			reporter = analyzer.NewReportStdoutTable()
		case "simple":
			reporter = analyzer.NewReportStdoutSimple()
		default:
			log.Fatalln("unknown format")
		}

		if err = analyzer.Analyze(dir, methods, reporter); err != nil {
			log.Fatalln(err.Error())
		}

		log.Println("analyze finished")
	},
}
