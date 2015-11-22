package main

import (
	"github.com/m0t0k1ch1/potto"
)

func PingAction(pctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	return potto.NewResponse("pong"), nil
}
