package analyzer

type isNewWork struct {
	name string
}

func NewisNew() isNewWork {
	return isNewWork{name: "New Wrok"}
}

func (c isNewWork) Name() string {
	return c.name
}

// checkPrice looks for cheaper jbos than before
func (c isNewWork) Method(data workMap) ([]AnaResult, error) {
	ret := []AnaResult{}
	for _, works := range data {
		if len(works) != 1 {
			continue
		}
		ret = append(ret, AnaResult{
			Report: "New Work",
			Work:   works[0],
		})
	}
	return ret, nil
}
