package http

import "github.com/rs/zerolog/log"

type Service struct {
}

func NewHttpService() *Service {
	return &Service{}
}

func (s *Service) Start() {
	log.Info().Msg("HTTP service started - To be implemented")
}
