package data

import (
	"thuanle/cse-mark/internal/models"
	"thuanle/cse-mark/internal/services/db"
	"time"
)

func TeacherStoreCourseLink(course string, link string, byId int64, byUser string) error {
	model := &models.CourseSettingsModel{
		Course:    course,
		Link:      link,
		ById:      byId,
		ByUser:    byUser,
		UpdatedAt: time.Now().Unix(),
	}
	return db.Instance().UpdateSubLinkSettings(model)
}

func IsTeacher(user string) bool {
	u, err := db.Instance().GetUserById(user)
	if err != nil {
		return false
	}
	return u.IsTeacher
}
