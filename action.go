package main

type Args []string
type Action func(*Context, Args) (*Response, error)

func NewActionMap() map[string]Action {
	actionMap := map[string]Action{}
	actionMap["ping"] = Ping

	return actionMap
}

func Ping(ctx *Context, args Args) (*Response, error) {
	return NewResponse("pong"), nil
}
