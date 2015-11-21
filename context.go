package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/m0t0k1ch1/ksatriya"
	"github.com/m0t0k1ch1/slackbot"
)

type Context struct {
	ksatriya.Ctx
	redisConn redis.Conn
	slackConn *slackbot.Client
	actions   map[string]Action
}

func convertContext(kctx ksatriya.Ctx) *Context {
	ctx, _ := kctx.(*Context)
	return ctx
}

func (ctx *Context) SetupActions() {
	ctx.actions["ping"] = Ping
}
