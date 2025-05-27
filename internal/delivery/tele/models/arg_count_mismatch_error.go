package models

import "fmt"

type ArgCountMismatchError struct {
	Needed int
	Actual int
}

func (a *ArgCountMismatchError) Error() string {
	return fmt.Sprintf("Arg count mismatch: needed %d, got %d", a.Needed, a.Actual)
}

func NewArgCountMismatchError(needed int, actual int) *ArgCountMismatchError {
	return &ArgCountMismatchError{
		Needed: needed,
		Actual: actual,
	}
}
