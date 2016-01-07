package main

import (
	"fmt"
	"strings"
)

func Show(ctx *Context, args ActionArgs) (*Response, error) {
	members, err := ctx.redis.SMEMBERS(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}

	var text string
	switch {
	case len(members) == 0:
		text = "no member"
	default:
		text = fmt.Sprintf("[ %s ]", strings.Join(members, ", "))
	}

	return response(text)
}
