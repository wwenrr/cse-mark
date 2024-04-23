package fetcher

import (
	"github.com/rs/zerolog/log"
	"thuanle/cse-mark/internal/data"
	"time"
)

func fetchNewMarks() {
	log.Info().Msg("Fetching new marks for all classes")

	courses, _ := data.GetAllCourses()

	for _, course := range courses {
		log.Info().Str("course", course.Course).Msg("Fetching new marks")
		msg, err := data.FetchCourseMarks(course.Course, course.Link)
		if err != nil {
			log.Error().
				Str("course", course.Course).
				Str("link", course.Link).
				Err(err).
				Msg("Fetch course marks error")
		}
		log.Info().
			Str("course", course.Course).
			Str("link", course.Link).
			Str("msg", msg).
			Msg("Fetched new marks")
		time.Sleep(time.Minute)
	}
}
