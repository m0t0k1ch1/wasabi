package main

import "github.com/m0t0k1ch1/potto"

type Response struct {
	*potto.Response
}

func NewResponse(text string) *Response {
	return &Response{potto.NewResponse(text)}
}
