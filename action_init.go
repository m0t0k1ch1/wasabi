package main

func Init(ctx *Context, args ActionArgs) (*Response, error) {
	if _, err := ctx.redis.DEL(ctx.ChannelID()); err != nil {
		return errorResponse(err)
	}
	return response("initialized")
}
