package main

import (
	"log"
	"os"

	"github.com/Nikkely/GLSite/internal/fetcher"
)

func main() {
	err := fetcher.Fetch()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("fetching finished")
	os.Exit(0)
}
