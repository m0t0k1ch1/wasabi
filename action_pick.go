package main

import "fmt"

func Pick(ctx *Context, args ActionArgs) (*Response, error) {
	len, err := ctx.Redis.SCARD(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}
	if len == 0 {
		return response("no member")
	}

	member, err := ctx.Redis.SRANDMEMBER(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}

	emoji := ctx.Conf.Emoji
	return response(fmt.Sprintf("%s %s %s", emoji, member, emoji))
}
