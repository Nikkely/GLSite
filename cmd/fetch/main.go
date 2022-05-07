package main

import (
	"flag"
	"log"
	"os"

	"github.com/Nikkely/GLSite/internal/fetcher"
)

func main() {
	dir := flag.String("path", "output", "output path")
	flag.Parse()

	err := fetcher.Fetch(*dir)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("fetching finished")
	os.Exit(0)
}
