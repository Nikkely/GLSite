package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/Nikkely/GLSite/model"
)

type Writer interface {
	Write(works []model.Work) error
}

type JSONWriter struct {
	OutputDir string
}

func (j JSONWriter) Write(works []model.Work) error {
	chs := make(chan error, len(works))
	wg := &sync.WaitGroup{}
	for _, work := range works {
		wg.Add(1)
		go func(w model.Work, c chan error) {
			defer wg.Done()
			raw, e := json.Marshal(w)
			if e != nil {
				c <- e
				return
			}
			c <- ioutil.WriteFile(joinFilePath(w.ID+"_"+w.FetchedAt.Format(time.RFC3339)+".json", j.OutputDir), raw, 0644)
		}(work, chs)
	}

	wg.Wait()
	close(chs)
	hasError := false
	for e := range chs {
		if e != nil {
			log.Println(e.Error())
			hasError = true
		}
	}
	if hasError {
		return fmt.Errorf("some error occured while writing")
	}
	return nil
}
