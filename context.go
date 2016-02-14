package main

import "github.com/m0t0k1ch1/potto"

type Context struct {
	*potto.Context
	Redis *Redis
	Conf  *Config
}

func (ctx *Context) ChannelID() string {
	return ctx.ParamSingle("channel_id")
}

func (ctx *Context) Finalize() {
	ctx.Redis.Close()
}
