package main

type Args []string
type Action func(*Context, Args) *Response

func NewActionMap() map[string]Action {
	var actionMap map[string]Action
	actionMap["ping"] = Ping

	return actionMap
}

func Ping(ctx *Context, args Args) *Response {
	return NewResponse(ctx.ParamSingle("channel_name"), "pong")
}
