package cmd

import (
	"log"

	"github.com/Nikkely/GLSite/internal/fetcher"
	"github.com/spf13/cobra"
)

const (
	fetchDirPath   = "dir-path"
	fetchFormat    = "format"
	fetchLimitPage = "limit-page"
)

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.PersistentFlags().StringP(fetchDirPath, "d", "output", "path to output direcotry")
	fetchCmd.PersistentFlags().StringP(fetchFormat, "f", "json", "specify format")
	fetchCmd.PersistentFlags().Int(fetchLimitPage, -1, "limit scraping pages for debug")
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Scrape and output",
	Long:  `Scrape and output. You can choose format and output path.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString(fetchDirPath)
		if err != nil {
			log.Fatalln(err.Error())
		}
		format, err := cmd.Flags().GetString(fetchFormat)
		if err != nil {
			log.Fatalln(err.Error())
		}
		limitPage, err := cmd.Flags().GetInt(fetchLimitPage)
		if err != nil {
			log.Fatalln(err.Error())
		}

		var options []fetcher.FetchOption
		if format == "json" {
			options = append(options, fetcher.OutputJSON(dir))
		}
		if limitPage != -1 {
			options = append(options, fetcher.LimitPage(limitPage))
		}

		if err := fetcher.Fetch(dir, options...); err != nil {
			log.Fatalln(err.Error())
		}

		log.Println("fetching finished")
	},
}
