package main

import "github.com/m0t0k1ch1/potto"

func Ping(ctx *Context, args potto.ActionArgs) (*potto.Response, error) {
	return response("pong")
}
