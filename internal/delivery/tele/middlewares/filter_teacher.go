package middlewares

import (
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/delivery/tele/models"
	"thuanle/cse-mark/internal/usecases/iam"
)

type TeacherOnly struct {
	authzService *iam.AuthzService
}

func NewTeacherOnly(authzService *iam.AuthzService) *TeacherOnly {
	return &TeacherOnly{
		authzService: authzService,
	}
}

func (m *TeacherOnly) Handle(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			defer c.Respond()
		}

		chatUsername := c.Chat().Username
		isTeacher, err := m.authzService.IsTeacher(chatUsername)
		if err == nil || isTeacher {
			return next(c)
		}

		return models.NewUnauthorizedError("you are not a teacher")
	}
}
