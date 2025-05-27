package coursequery

import (
	"thuanle/cse-mark/internal/domain/course"
	"time"
)

type ActiveCourseService struct {
	CourseRepo course.Repository
	Rule       *course.Rules
	Now        func() time.Time
}

func NewActiveCourseService(r course.Repository, rule *course.Rules) *ActiveCourseService {
	return &ActiveCourseService{
		CourseRepo: r,
		Rule:       rule,
		Now:        time.Now,
	}
}

func (s *ActiveCourseService) ListActiveCourses() ([]course.Model, error) {
	threshold := s.Now().Add(-s.Rule.CourseActiveAge)

	actives, err := s.CourseRepo.FindCoursesUpdatedAfter(threshold)
	if err != nil {
		return nil, err
	}

	return actives, nil
}
