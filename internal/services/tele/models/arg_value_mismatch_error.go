package models

import "fmt"

type ArgValueMismatchError struct {
	msg string
}

func (a *ArgValueMismatchError) Error() string {
	return fmt.Sprintf("Arg value mismatch: %s", a.msg)
}

func NewArgValueMismatchError(msg string) *ArgValueMismatchError {
	return &ArgValueMismatchError{
		msg: msg,
	}
}
