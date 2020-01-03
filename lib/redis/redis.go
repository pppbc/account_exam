package redis11

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	max_idle     = 30
	max_active   = 30
	idle_timeout = time.Duration(60)
)

var RedisPool = &redis.Pool{
	MaxIdle:     max_idle,
	MaxActive:   max_active,
	IdleTimeout: idle_timeout,
	Wait:        true,
	Dial: func() (redis.Conn, error) {
		con, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			return nil, err
		}
		return con, nil
	},
}
