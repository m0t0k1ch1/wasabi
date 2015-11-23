package main

import "github.com/m0t0k1ch1/potto"

func Ping(pctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	return response("pong")
}
