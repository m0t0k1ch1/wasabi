package main

import (
	"github.com/m0t0k1ch1/potto"
)

func PingAction(pctx potto.Ctx, args potto.Args) (*Response, error) {
	ctx, _ := convertContext(pctx)
	return c.Ping(ctx)
}
func Ping(ctx *Context) {
	return potto.NewResponse("pong"), nil
}
