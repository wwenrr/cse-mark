package handlers

import (
	"net/url"
	"thuanle/cse-mark/internal/delivery/tele/handlers/helpers"
	"thuanle/cse-mark/internal/delivery/tele/models"
	"thuanle/cse-mark/internal/delivery/tele/views"
	"thuanle/cse-mark/internal/domain/course"
	"thuanle/cse-mark/internal/domain/mark"
	"thuanle/cse-mark/internal/usecases/iam"
	"thuanle/cse-mark/internal/usecases/markimport"

	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

type Teacher struct {
	courseRepo  course.Repository
	courseRules *course.Rules

	teacherViewRender *views.TeacherRenderer

	authzService *iam.AuthzService

	markRepo          mark.Repository
	markImportService *markimport.Service
}

func NewTeacherHandler(courseRepo course.Repository, courseRules *course.Rules,
	teacherViewRender *views.TeacherRenderer,
	authzService *iam.AuthzService,
	markRepo mark.Repository,
	markImportService *markimport.Service) *Teacher {
	return &Teacher{
		courseRepo:  courseRepo,
		courseRules: courseRules,

		teacherViewRender: teacherViewRender,

		authzService: authzService,

		markImportService: markImportService,
		markRepo:          markRepo,
	}
}

func (h *Teacher) GetMyProfile(c telebot.Context) error {
	chatUsername := c.Chat().Username

	log.Info().
		Str("chatUsername", chatUsername).
		Msg("Get teacher profile")

	courses, err := h.courseRepo.FindCoursesManagedByUser(chatUsername)
	if err != nil {
		return err
	}

	msg := h.teacherViewRender.RenderTeacherProfile(courses)
	return helpers.SendPre(c, msg)
}

func (h *Teacher) LoadCourseLink(c telebot.Context) error {
	courseId, link, err := helpers.Args2StrStr(c)
	if err != nil {
		return err
	}

	if !h.courseRules.IsValidCourseId(courseId) {
		return models.NewArgValueMismatchError("courseId invalid")
	}

	_, err = url.ParseRequestURI(link)
	if err != nil {
		return err
	}

	chatId := c.Chat().ID
	chatUsername := c.Chat().Username

	isGranted, err := h.authzService.CanEditCourse(chatUsername, chatId, courseId)
	if err != nil || !isGranted {
		return models.NewUnauthorizedError("cannot modify courseId")
	}

	log.Info().
		Int64("chatId", chatId).
		Str("chatUsername", chatUsername).
		Str("courseId", courseId).
		Str("link", link).
		Msg("Admin store marks")

	err = h.courseRepo.UpdateCourseLink(courseId, link, chatId, chatUsername)
	if err != nil {
		return err
	}

	count, err := h.markImportService.FetchMarkLinkIntoCourse(courseId, link)
	if err != nil {
		return err
	}

	return helpers.Sendf(c, "%s: Store %d records.", courseId, count)
}

func (h *Teacher) ClearCourseLink(c telebot.Context) error {
	courseId, err := helpers.Args2Str(c)
	if err != nil {
		return err
	}

	if !h.courseRules.IsValidCourseId(courseId) {
		return models.NewArgValueMismatchError("course invalid")
	}

	chatId := c.Chat().ID
	chatUsername := c.Chat().Username

	isGranted, err := h.authzService.CanEditCourse(chatUsername, chatId, courseId)
	if err != nil || !isGranted {
		return models.NewUnauthorizedError("cannot modify courseId")
	}

	log.Info().
		Str("courseId", courseId).
		Str("chatUsername", chatUsername).
		Int64("chatId", chatId).
		Msg("Clear marks")

	err = h.markRepo.RemoveCourseMarks(courseId)
	if err != nil {
		return err
	}

	err = h.courseRepo.RemoveCourse(courseId)

	return helpers.Sendf(c, "%s: cleared", courseId)
}
