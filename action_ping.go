package main

func Ping(ctx *Context, args ActionArgs) (*Response, error) {
	return response("pong")
}
