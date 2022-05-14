package cmd

import (
	"log"

	"github.com/Nikkely/GLSite/analyzer"
	"github.com/spf13/cobra"
)

const (
	analyzeDirPathFlg     = "dir-path"
	analyzeDLThresholdFlg = "dl-th"
)

func init() {
	rootCmd.AddCommand(AnalCmd)

	AnalCmd.PersistentFlags().StringP(analyzeDirPathFlg, "d", "output", "path to output direcotry")
	AnalCmd.PersistentFlags().Int(analyzeDLThresholdFlg, 50, "path to output direcotry")
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

		methods := []analyzer.AnalyzeMethod{
			analyzer.NewChangePrice(),
			analyzer.NewisNew(),
			analyzer.NewChangeDL(dlth),
		}

		if err = analyzer.Analyze(dir, methods...); err != nil {
			log.Fatalln(err.Error())
		}

		log.Println("analyze finished")
	},
}
