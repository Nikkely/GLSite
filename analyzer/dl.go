package analyzer

import (
	"fmt"
	"log"
	"sort"
)

type changeDL struct {
	name      string
	threshold int
}

func NewChangeDL(threshold int) changeDL {
	return changeDL{name: "DL数UP", threshold: threshold}
}

func (c changeDL) Name() string {
	return c.name
}

type cahngeDLResult struct {
	diff   int
	result AnaResult
}

// checkPrice looks for cheaper jbos than before
func (c changeDL) Method(data workMap) ([]AnaResult, error) {
	res := []cahngeDLResult{}
	log.Printf("Threshold: %d", c.threshold)
	for _, works := range data {
		if len(works) < 2 {
			continue
		}

		now := works[0]
		pre := works[1]
		count := now.DL - pre.DL
		if count < c.threshold {
			continue
		}
		res = append(res, cahngeDLResult{
			diff: count,
			result: AnaResult{
				Report: fmt.Sprintf("ダウンロード数が%d件増えました", count),
				Work:   now,
			},
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].diff > res[j].diff
	})
	ret := []AnaResult{}
	for _, r := range res {
		ret = append(ret, r.result)
	}
	return ret, nil
}
