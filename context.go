package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/m0t0k1ch1/potto"
)

type Context struct {
	potto.Ctx
	conf      *Config
	redisConn redis.Conn
}

func convertContext(pctx potto.Ctx) *Context {
	ctx, _ := pctx.(*Context)
	return ctx
}
