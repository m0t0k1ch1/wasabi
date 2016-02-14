package main

import "fmt"

func Add(ctx *Context, args ActionArgs) (*Response, error) {
	channelID := ctx.ChannelID()

	totalCount := 0
	var lastAddedMember string
	for _, arg := range args {
		count, err := ctx.Redis.SADD(channelID, arg)
		if err != nil {
			return errorResponse(err)
		}
		totalCount += count
		lastAddedMember = arg
	}

	var text string
	switch {
	case totalCount == 0:
		text = "already exist"
	case totalCount > 1:
		text = fmt.Sprintf("added %d members", totalCount)
	default:
		text = fmt.Sprintf("added %s", lastAddedMember)
	}

	return response(text)
}
