package mark

import "errors"

var ErrNotFound = errors.New("no marks in result")

type Repository interface {
	GetMark(courseId string, studentId string) (string, error)

	RemoveMarksByCourseId(courseId string) error
	AddCourseMarks(courseId string, marks []map[string]string) error
	RemoveCourseMarks(courseId string) error
}
