package testable

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	connstr string
	err     error
}

func RedisFrom(v string) *Redis {
	return &Redis{connstr: v}
}

func (r *Redis) Error() error {
	return r.err
}

func (r *Redis) Name() string {
	return "Redis"
}

func (r *Redis) Test() error {
	if r.err != nil {
		return r.Error()
	}

	c, err := redis.DialURL(r.connstr)
	if err != nil {
		r.err = err
		return r.Error()
	}
	defer c.Close()

	pong, err := c.Do("PING")
	if err != nil {
		r.err = err
		return r.Error()
	}
	if pong != "PONG" {
		r.err = fmt.Errorf("Should receive 'PONG' but was '%s'", pong)
		return r.Error()
	}

	return nil
}
