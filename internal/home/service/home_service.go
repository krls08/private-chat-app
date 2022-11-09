package service

import (
	"fmt"
	"io/ioutil"
	"log"
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
	s.GetStaticFiles()
	return template, nil
}

func (s *DefaultHomeService) GetStaticFiles() {
	files, err := ioutil.ReadDir("./static/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
