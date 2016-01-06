package main

import (
	"fmt"

	"github.com/m0t0k1ch1/potto"
)

func Pick(ctx *Context, args potto.ActionArgs) (*potto.Response, error) {
	len, err := ctx.redis.SCARD(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}
	if len == 0 {
		return response("no member")
	}

	member, err := ctx.redis.SRANDMEMBER(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}

	emoji := ctx.conf.Emoji
	return response(fmt.Sprintf("%s %s %s", emoji, member, emoji))
}
