package client

import "errors"

var (
	ErrorsServerError                 = errors.New("server error")
	ErrorMessageTypeProcessorIsExists = errors.New("this message type processor is Exists")
)
