package testable

import (
	"fmt"
	"strings"
	"sync"
)

func env2testable(e string, ch chan Testable, wg *sync.WaitGroup) {
	defer wg.Done()

	if !strings.HasPrefix(e, "TEST_") {
		// Not a test env
		return
	}

	// Split environ as key value pair
	k, v := func() (string, string) {
		x := strings.SplitN(e, "=", 2)
		return x[0], x[1]
	}()

	k = strings.TrimPrefix(k, "TEST_")

	if strings.HasPrefix(k, "SQL_") {
		ch <- SQLFrom(strings.TrimPrefix(k, "SQL_"), v)
	} else if strings.HasPrefix(k, "REDIS") {
		ch <- RedisFrom(v)
	} else if strings.HasPrefix(k, "HTTP_") {
		ch <- HTTPFrom(v)
	} else {
		ch <- &NotImplemented{err: fmt.Errorf("Test type %s is not implemented", k)}
	}
}

func FromEnv(env []string) []Testable {
	var wg sync.WaitGroup
	wg.Add(len(env))

	ch := make(chan Testable, len(env))
	for _, e := range env {
		go env2testable(e, ch, &wg)
	}
	wg.Wait()
	close(ch)

	r := make([]Testable, 0)
	for t := range ch {
		r = append(r, t)
	}
	return r
}
