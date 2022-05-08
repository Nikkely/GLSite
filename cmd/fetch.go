package cmd

import (
	"log"

	"github.com/Nikkely/GLSite/internal/fetcher"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Scrape and output",
	Long:  `Scrape and output. You can choose format and specify output path.`,
	Run: func(cmd *cobra.Command, args []string) {
		var format, dir string
		cmd.Flags().StringVarP(&dir, "dir-path", "d", "output", "path to output direcotry")
		cmd.Flags().StringVarP(&format, "format", "f", "json", "specify format")
		switch format {
		case "json":
			j := fetcher.NewJSONWriter(dir)
			err := fetcher.Fetch(j)
			if err != nil {
				log.Fatalln(err.Error())
			}
		default:
			log.Fatalln("Unknown format")
		}
		log.Println("fetching finished")
	},
}
