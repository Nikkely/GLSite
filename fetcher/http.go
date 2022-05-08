package fetcher

import "fmt"

type notFoundErr struct {
	err error
}
type otherErr struct {
	err error
}

func newNotFoundErr() notFoundErr {
	return notFoundErr{err: fmt.Errorf("404 Not Found")}
}

func (e notFoundErr) Error() string {
	return e.err.Error()
}

func newOtherErr(status string) otherErr {
	return otherErr{err: fmt.Errorf(status)}
}

func (e otherErr) Error() string {
	return e.err.Error()
}
