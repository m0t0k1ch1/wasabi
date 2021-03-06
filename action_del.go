package main

import "fmt"

func Del(ctx *Context, args ActionArgs) (*Response, error) {
	channelID := ctx.ChannelID()

	totalCount := 0
	var lastDeletedMember string
	for _, arg := range args {
		count, err := ctx.Redis.SREM(channelID, arg)
		if err != nil {
			return errorResponse(err)
		}
		totalCount += count
		lastDeletedMember = arg
	}

	var text string
	switch {
	case totalCount == 0:
		text = "not exist"
	case totalCount > 1:
		text = fmt.Sprintf("deleted %d members", totalCount)
	default:
		text = fmt.Sprintf("deleted %s", lastDeletedMember)
	}

	return response(text)
}
