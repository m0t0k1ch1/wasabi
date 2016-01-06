package main

import "github.com/m0t0k1ch1/potto"

func Init(ctx *Context, args potto.ActionArgs) (*potto.Response, error) {
	if _, err := ctx.redis.DEL(ctx.ChannelID()); err != nil {
		return errorResponse(err)
	}
	return response("initialized")
}
