package main

func response(text string) (*Response, error) {
	return NewResponse(text), nil
}

func errorResponse(err error) (*Response, error) {
	return NewResponse(err.Error()), err
}
