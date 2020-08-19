package testable

import (
	"fmt"
	"net/http"
)

type HTTP struct {
	url  string
	name string
	err  error
}

func HTTPFrom(value string) *HTTP {
	return &HTTP{
		url: value,
	}
}

func (h *HTTP) Error() error {
	return h.err
}

func (h *HTTP) Name() string {
	return "Http " + h.url
}

func (h *HTTP) Test() error {
	if h.err != nil {
		return h.Error()
	}

	resp, err := http.Head(h.url)
	if err != nil {
		h.err = err
		return h.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		h.err = fmt.Errorf(resp.Status)
		return h.Error()
	}

	return h.Error()
}
