package main

func Init(ctx *Context, args ActionArgs) (*Response, error) {
	if _, err := ctx.Redis.DEL(ctx.ChannelID()); err != nil {
		return errorResponse(err)
	}
	return response("initialized")
}
