package analyzer

import (
	"fmt"
	"sort"
)

type changeDL struct {
	name      string
	threshold int
}

func NewChangeDL(threshold int) changeDL {
	return changeDL{name: "DLCount", threshold: threshold}
}

func (c changeDL) Name() string {
	return c.name
}

// checkPrice looks for cheaper jbos than before
func (c changeDL) Method(data workMap) ([]AnaResult, error) {
	ret := []AnaResult{}
	for _, works := range data {
		if len(works) < 2 {
			continue
		}

		now := works[0]
		pre := works[1]
		count := now.DL-pre.DL
		if  count < c.threshold {
			continue
		}
		ret = append(ret, AnaResult{
			Report: fmt.Sprintf(`download count add %d (>%d)`, count, c.threshold),
			Work:   now,
		})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Work.DL > ret[j].Work.DL
	})
	return ret, nil
}
