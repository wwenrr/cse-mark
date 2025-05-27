package user

import "errors"

var ErrNotFound = errors.New("no users in result")

type Repository interface {
	UpdateUser(username string, isTeacher bool, grantedBy string) error
	FindUserById(username string) (Model, error)
}
