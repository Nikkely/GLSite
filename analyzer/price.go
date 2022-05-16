package analyzer

import (
	"fmt"
	"sort"
)

type changePrice struct {
	name string
}

func NewChangePrice() changePrice {
	return changePrice{name: "割引作品"}
}

func (c changePrice) Name() string {
	return c.name
}

// checkPrice looks for cheaper jbos than before
func (c changePrice) Method(data workMap) ([]AnaResult, error) {
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
			Report: fmt.Sprintf(`%d 円割引中です`, discounted),
			Work:   now,
		})
	}
	sort.Slice(ret, func(i, j int) bool {
		iPrice := ret[i].Work.Discount
		if iPrice == 0 {
			iPrice = ret[i].Work.Price
		}
		jPrice := ret[j].Work.Discount
		if jPrice == 0 {
			jPrice = ret[j].Work.Price
		}
		return iPrice > jPrice
	})
	return ret, nil
}
