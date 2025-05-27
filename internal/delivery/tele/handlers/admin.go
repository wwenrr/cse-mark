package handlers

import (
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/delivery/tele/handlers/helpers"
	"thuanle/cse-mark/internal/domain/teleuser"
	"thuanle/cse-mark/internal/domain/user"
)

type Admin struct {
	userRepo user.Repository
}

func NewAdminHandler(userRepo user.Repository) *Admin {
	return &Admin{
		userRepo: userRepo,
	}
}

func (h *Admin) SetTeacher(c telebot.Context) error {
	name, grant, err := helpers.Args2StrDbool(c, true)
	if err != nil {
		return err
	}

	if !teleuser.IsValidTelegramUsername(name) {
		return helpers.SendErrorArgumentValueMismatch(c, "name incorrect")
	}

	chatName := c.Chat().Username

	err = h.userRepo.UpdateUser(name, grant, chatName)
	if err != nil {
		return err
	}

	if grant {
		return helpers.Sendf(c, "Set %s as teacher", name)
	} else {
		return helpers.Sendf(c, "Remove %s from teacher", name)
	}
}
