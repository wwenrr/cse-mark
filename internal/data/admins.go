package data

import (
	"thuanle/cse-mark/internal/models"
	"thuanle/cse-mark/internal/services/db"
)

func AdminGrantTeacher(user string, grant bool, grantedBy string) error {
	uModel := &models.UserSettingsModel{
		UserId:    user,
		IsTeacher: grant,
		GrantedBy: grantedBy,
	}
	return db.Instance().UpdateUserSettings(uModel)
}
