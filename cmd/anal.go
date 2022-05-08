package cmd

import (
	"log"

	"github.com/Nikkely/GLSite/analyzer"
	"github.com/spf13/cobra"
)

const analyzeDirPath = "dir-path"

func init() {
	rootCmd.AddCommand(AnalCmd)

	AnalCmd.PersistentFlags().StringP(analyzeDirPath, "d", "output", "path to output direcotry")
}

var AnalCmd = &cobra.Command{
	Use:   "anal",
	Short: "Analyzer",
	Long:  `Analyze datas`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString(fetchDirPath)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if err = analyzer.Analyze(dir); err != nil {
			log.Fatalln(err.Error())
		}

		log.Println("analyze finished")
	},
}
