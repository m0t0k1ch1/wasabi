package main

import (
	"strings"

	"github.com/m0t0k1ch1/potto"
)

func Show(pctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	return show(pctx.(*Context))
}
func show(ctx *Context) (*potto.Response, error) {
	members, err := ctx.redis.SMEMBERS(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}
	return response(strings.Join(members, ", "))
}
