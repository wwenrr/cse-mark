package marksync

import (
	"github.com/rs/zerolog/log"
	"thuanle/cse-mark/internal/domain/downloader"
	"thuanle/cse-mark/internal/usecases/coursequery"
	"thuanle/cse-mark/internal/usecases/markimport"
	"time"
)

type Service struct {
	courseQueryService *coursequery.ActiveCourseService

	downloader        downloader.Repository
	markImportService *markimport.Service

	fetchingInterval time.Duration
}

func NewService(courseQueryService *coursequery.ActiveCourseService,
	downloader downloader.Repository, markImportService *markimport.Service) *Service {
	return &Service{
		courseQueryService: courseQueryService,

		downloader:        downloader,
		markImportService: markImportService,

		fetchingInterval: time.Minute,
	}
}

func (s *Service) fetchNewMarks() {
	log.Info().Msg("Fetching new marks for all classes")

	courses, err := s.courseQueryService.ListActiveCourses()

	if err != nil {
		log.Error().Err(err).Msg("Fetching active courses error")
		return
	}

	for _, course := range courses {
		_, _ = s.markImportService.FetchMarkLinkIntoCourse(course.Id, course.Link)

		time.Sleep(s.fetchingInterval)
	}
}

func (s *Service) Run() {
	for {
		s.fetchNewMarks()
		log.Info().Msg("Sleeping for 10 minutes...")
		time.Sleep(10 * time.Minute)
	}
}
