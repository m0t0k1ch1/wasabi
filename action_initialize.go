package main

import "github.com/m0t0k1ch1/potto"

func Initialize(pctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	return initialize(pctx.(*Context))
}
func initialize(ctx *Context) (*potto.Response, error) {
	if _, err := ctx.redis.DEL(ctx.ChannelID()); err != nil {
		return errorResponse(err)
	}
	return response("initialized")
}
