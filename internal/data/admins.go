package data

import (
	"thuanle/cse-mark/internal/domain/entities"
	"thuanle/cse-mark/internal/services/db"
)

func AdminGrantTeacher(user string, grant bool, grantedBy string) error {
	uModel := &entities.UserSettingsModel{
		UserId:    user,
		IsTeacher: grant,
		GrantedBy: grantedBy,
	}
	return db.Instance().UpdateUserSettings(uModel)
}
