package models

import "fmt"

type UnauthorizedError struct {
	msg string
}

func (a *UnauthorizedError) Error() string {
	return fmt.Sprintf("Unauthorized action: %s", a.msg)
}

func NewUnauthorizedError(msg string) *UnauthorizedError {
	return &UnauthorizedError{
		msg: msg,
	}
}
