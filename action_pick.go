package main

import (
	"fmt"

	"github.com/m0t0k1ch1/potto"
)

func Pick(pctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	return pick(pctx.(*Context))
}
func pick(ctx *Context) (*potto.Response, error) {
	member, err := ctx.redis.SRANDMEMBER(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}

	emoji := ctx.conf.Emoji
	return response(fmt.Sprintf("%s %s %s", emoji, member, emoji))
}
