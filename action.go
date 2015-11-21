package main

import (
	"fmt"

	"github.com/m0t0k1ch1/ksatriya"
)

func PingHandler(kctx ksatriya.Ctx) {
	ping(convertContext(kctx))
}
func ping(ctx *Context) {
	channel := fmt.Sprintf("#%s", ctx.ParamSingle("channel_name"))
	ctx.slackConn.SendMessage(channel, "pong")
}
