package analyzer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Nikkely/GLSite/model"
)

// workMap is a map. key is model.work.ID and value is []model.work sorted by model.work.fetchedAt desc.
type workMap map[string][]model.Work

// AnaResult is result of analyze
type AnaResult struct {
	Work   model.Work
	Report string
}

type AnalyzeMethod interface {
	Name() string
	Method(data workMap) ([]AnaResult, error)
}

func Analyze(dir string) error {
	wm, err := load(dir)
	if err != nil {
		return err
	}

	analyzers := []AnalyzeMethod{
		NewChangePrice(),
	}
	for _, a := range analyzers {
		log.Printf(`Analyze: %s`, a.Name())
		var res []AnaResult
		if res, err = a.Method(wm); err != nil {
			return err
		}
		if len(res) == 0 {
			log.Printf("No result")
		} else {
			ReportStdout(res)
		}
	}
	return nil
}

// load loads works from dir
func load(dir string) (workMap, error) {
	ret := workMap{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		if err != nil {
			return err
		}

		var raw []byte
		if raw, err = os.ReadFile(path); err != nil {
			return err
		}
		var work model.Work
		if err = json.Unmarshal(raw, &work); err != nil {
			return fmt.Errorf("failed loading %s: %s", path, err.Error())
		}

		if _, ok := ret[work.ID]; !ok {
			ret[work.ID] = []model.Work{work}
		} else {
			ret[work.ID] = append(ret[work.ID], work)
		}
		return nil
	})
	for _, v := range ret {
		sort.Slice(v, func(i, j int) bool {
			return v[i].FetchedAt.After(v[j].FetchedAt)
		})
	}
	return ret, err
}

// checkPrice looks for cheaper jbos than before
func checkPrice(data workMap) ([]AnaResult, error) {
	ret := []AnaResult{}
	for _, works := range data {
		if len(works) < 2 {
			continue
		}

		now := works[0]
		pre := works[1]
		discounted := 0
		if now.Discount == 0 && pre.Discount == 0 && now.Price < pre.Price {
			discounted = pre.Price - now.Price
		} else if now.Discount != 0 && pre.Discount == 0 {
			discounted = pre.Price - now.Discount
		} else if now.Discount != 0 && pre.Discount != 0 && now.Discount < pre.Discount {
			discounted = pre.Discount - now.Discount
		}
		if discounted == 0 {
			continue
		}
		ret = append(ret, AnaResult{
			Report: fmt.Sprintf(`Dicounted %d.`, discounted),
			Work:   now,
		})
	}
	return ret, nil
}
