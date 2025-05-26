package data

import (
	"thuanle/cse-mark/internal/domain/entities"
	"thuanle/cse-mark/internal/services/db"
)

func TeacherStoreCourseLink(course string, link string, byId int64, byUser string) error {
	model := &entities.CourseSettingsModel{
		Course: course,
		Link:   link,
		ById:   byId,
		ByUser: byUser,
	}
	return db.Instance().SetCourse(model)
}

func IsTeacher(user string) bool {
	u, err := db.Instance().GetUserById(user)
	if err != nil {
		return false
	}
	return u.IsTeacher
}

func FetchTeacherCourses(user string) ([]*entities.CourseSettingsModel, error) {
	return db.Instance().GetCoursesByUser(user)
}
