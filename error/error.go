package main

import "fmt"

type error interface {
	Error() string
}

type CustomError struct {
	error         error
	errorCode     ErrorCode
	errorMsg      string
	loggingParams map[string]interface{}
	exists        bool
}

type ErrorCode string

func (e ErrorCode) ToString() string {
	return fmt.Sprintf("%s", e)
}

const (
	SQL_INSERT_ERROR ErrorCode = "SQL_INSERT_ERROR"
	SQL_UPDATE_ERROR ErrorCode = "SQL_UPDATE_ERROR"
)

var ErrorCodeToUserMessages = map[ErrorCode]string {
	REQUEST_NOT_VALID: "Invalid request. Please try again",
	SQL_INSERT_ERROR: "Something went wrong, please try again",
	SQL_UPDATE_ERROR: "Something went wrong, please try again"
}
