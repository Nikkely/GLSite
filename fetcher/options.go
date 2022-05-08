package fetcher

type fetchOptions struct {
	writer      writer
	limitPage   int
	unlimitPage bool
}
type FetchOption func(*fetchOptions)

func defaultOption(dir string) fetchOptions {
	return fetchOptions{
		writer:      jsonWriter{outputDir: dir},
		limitPage:   0,
		unlimitPage: true,
	}
}

func OutputJSON(dir string) FetchOption {
	return func(o *fetchOptions) {
		o.writer = jsonWriter{outputDir: dir}
	}
}

func LimitPage(p int) FetchOption {
	return func(o *fetchOptions) {
		o.limitPage = p
		o.unlimitPage = false
	}
}
