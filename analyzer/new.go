package analyzer

type isNewWork struct {
	name string
}

func NewisNew() isNewWork {
	return isNewWork{name: "新作"}
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
			Report: "",
			Work:   works[0],
		})
	}
	return ret, nil
}
