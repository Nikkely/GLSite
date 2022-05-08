package fetcher

type fetchOptions struct {
	writer      Writer
	limitPage   int
	unlimitPage bool
}
type FetchOption func(*fetchOptions)

func defaultOption(dir string) fetchOptions {
	return fetchOptions{
		writer:      JSONWriter{OutputDir: dir},
		limitPage:   0,
		unlimitPage: true,
	}
}

func OutputJSON(dir string) FetchOption {
	return func(o *fetchOptions) {
		o.writer = JSONWriter{OutputDir: dir}
	}
}

func LimitPage(p int) FetchOption {
	return func(o *fetchOptions) {
		o.limitPage = p
		o.unlimitPage = false
	}
}
