package main

import (
	"fmt"
	"strings"

	"github.com/m0t0k1ch1/potto"
)

func Show(pctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	ctx := pctx.(*Context)

	members, err := ctx.redis.SMEMBERS(ctx.ChannelID())
	if err != nil {
		return errorResponse(err)
	}

	var text string
	switch {
	case len(members) == 0:
		text = "no member"
	default:
		text = fmt.Sprintf("targets: %s", strings.Join(members, ", "))
	}

	return response(text)
}
