package course

import (
	"github.com/rs/zerolog/log"
	"regexp"
	"slices"
	"thuanle/cse-mark/internal/configs"
	"time"
)

type Rules struct {
	CourseActiveAge time.Duration

	adminIds []int64
}

func NewRules(config *configs.Config) *Rules {
	return &Rules{
		CourseActiveAge: config.CourseActiveAge,
		adminIds:        config.TeleAdminChatIds,
	}
}

// IsCourseActive checks if the course is active based on the last updated time.
func (r *Rules) IsCourseActive(course Model) bool {
	return time.Since(time.Unix(course.UpdatedAt, 0)) < r.CourseActiveAge
}

// IsValidCourseId checks if the course ID is valid.
func (r *Rules) IsValidCourseId(id string) bool {
	pattern := "^[a-zA-Z][a-zA-Z0-9-]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(id)
}

// CourseUpdateTill calculates the time until which the course is considered active based on its last updated time.
func (r *Rules) CourseUpdateTill(course Model) time.Time {
	return time.Unix(course.UpdatedAt, 0).Add(r.CourseActiveAge)
}

// CanUserEditCourse checks if a user can edit a course based on the course's ownership and activity status.
func (r *Rules) CanUserEditCourse(course Model, user string, chatId int64) bool {
	if slices.Contains(r.adminIds, chatId) {
		log.Debug().
			Str("user", user).
			Int64("chatId", chatId).
			Ints64("adminChatIds", r.adminIds).
			Msg("Admin can modify course")
		return true
	}

	return course.ByTeleUser == user || course.ByTeleId == chatId
}
