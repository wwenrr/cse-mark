package data

import (
	"thuanle/cse-mark/internal/models"
	"thuanle/cse-mark/internal/services/db"
)

func TeacherStoreCourseLink(course string, link string, byId int64, byUser string) error {
	model := &models.CourseSettingsModel{
		Course: course,
		Link:   link,
		ById:   byId,
		ByUser: byUser,
	}
	return db.Instance().SetCourseSettings(model)
}

func IsTeacher(user string) bool {
	u, err := db.Instance().GetUserById(user)
	if err != nil {
		return false
	}
	return u.IsTeacher
}
