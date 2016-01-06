package main

import (
	"fmt"
	"strings"

	"github.com/m0t0k1ch1/potto"
)

func Show(ctx *Context, args potto.ActionArgs) (*potto.Response, error) {
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
