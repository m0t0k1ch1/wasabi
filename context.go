package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/m0t0k1ch1/potto"
)

type Context struct {
	*potto.Context
	conf      *Config
	redisConn redis.Conn
}
