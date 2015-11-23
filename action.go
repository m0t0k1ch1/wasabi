package main

import "github.com/m0t0k1ch1/potto"

func response(text string) (*potto.Response, error) {
	return potto.NewResponse(text), nil
}

func errorResponse(err error) (*potto.Response, error) {
	return potto.NewResponse(err.Error()), err
}
