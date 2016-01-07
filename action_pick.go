package main

import "fmt"

func Pick(ctx *Context, args ActionArgs) (*Response, error) {
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
