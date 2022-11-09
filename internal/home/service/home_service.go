package service

import (
	"strings"

	"github.com/krls08/private-chat-app/internal/home/domain"
)

type HomeService interface {
	GetTemplate(ipPort string) (string, error)
}

type DefaultHomeService struct {
	//home  domain.Home
	homer domain.Homer
}

func NewHomeService() *DefaultHomeService {
	return &DefaultHomeService{
		//homer: new(domain.Home),
		homer: &domain.Home{},
	}
}

func (s *DefaultHomeService) GetTemplate(ipPort string) (string, error) {
	ip := strings.Split(ipPort, ":")
	template := s.homer.FindTemplate(ip[0])
	return template, nil
}
