package main

type Action func(*Context, []string) *Response

func NewActionMap() map[string]Action {
	var actionMap map[string]Action
	actionMap["ping"] = Ping

	return actionMap
}

func Ping(ctx *Context, args []string) *Response {
	return NewResponse(ctx.ParamSingle("channel_name"), "pong")
}
