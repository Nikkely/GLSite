package main

import (
	"flag"
	"log"
	"os"

	"github.com/Nikkely/GLSite/internal/fetcher"
)

func main() {
	format := flag.String("format", "json", "specify format")
	dir := flag.String("path", "output", "output path")
	flag.Parse()

	switch *format {
	case "json":
		j := fetcher.NewJSONWriter(*dir)
		err := fetcher.Fetch(j)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	log.Println("fetching finished")
	os.Exit(0)
}
