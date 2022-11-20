package errors

import "net/http"

type MyError struct {
	Code int
	Msg  string
}

func Success(message string) *MyError {
	return &MyError{
		Code: http.StatusOK,
		Msg:  message,
	}
}

func Failed(code int, msg string) *MyError {
	return &MyError{code, msg}
}
