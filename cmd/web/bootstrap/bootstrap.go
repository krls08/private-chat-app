package bootstrap

import (
	"context"
	"log"

	"github.com/krls08/private-chat-app/internal/home/service"
	"github.com/krls08/private-chat-app/internal/infrastructure/server"
	"github.com/krls08/private-chat-app/internal/infrastructure/server/handlers"
)

func Run() error {

	hs := service.NewHomeService()
	hh := handlers.HomeHandlers{
		Service: hs,
	}

	go handlers.ListenToWsChannel()

	log.Println("Starting webserver on port 60002")
	ctx := context.TODO()
	s := server.New(ctx, "localhost", 60002, hh)

	return s.Run(ctx)
}

//func Run() error {
//
//	hs := service.NewHomeService()
//	hh := handlers.HomeHandlers{
//		Service: hs,
//	}
//
//	mux := routes.Mux(hh)
//
//	go handlers.ListenToWsChannel()
//
//	log.Println("Starting webserver on port 60002")
//
//	err := http.ListenAndServe(":60002", mux)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func Run() error {
//
//	hs := service.NewHomeService()
//	hh := handlers.HomeHandlers{
//		Service: hs,
//	}
//
//	mux := routes.Mux(hh)
//
//	go handlers.ListenToWsChannel()
//
//	log.Println("Starting webserver on port 60002")
//
//	err := http.ListenAndServe(":60002", mux)
//	if err != nil {
//		return err
//	}
//	return nil
//}
