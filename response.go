package main

import "fmt"

type Response struct {
	Channel string
	Text    string
}

func NewResponse(channel, text string) *Response {
	return &Response{
		Channel: fmt.Sprintf("#%s", channel),
		Text:    text,
	}
}
