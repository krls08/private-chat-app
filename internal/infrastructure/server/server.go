package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/krls08/private-chat-app/internal/home/service"
	"github.com/krls08/private-chat-app/internal/infrastructure/server/handlers"
)

type Server struct {
	httpAddr string
	mux      *http.ServeMux

	hh handlers.HomeHandlers
	hs service.HomeService
}

func New(ctx context.Context, host string, port uint, homeHandlers handlers.HomeHandlers) Server { //(context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		mux:      http.NewServeMux(),

		// use cases
		hh: homeHandlers,
	}

	srv.registerRoutes()
	//return serverContext(ctx), srv
	return srv

}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Listening on %s\n", s.httpAddr)

	return http.ListenAndServe(s.httpAddr, s.mux)
}

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("/", s.hh.Home_n)
	s.mux.HandleFunc("/ws", handlers.WsEndpoint)
	//s.mux.GET("/", s.hh.Home_g())
	//s.engine.GET("/ws", handlers.WsEndpoint())

	// Serve static files from /static folder
	fs := http.FileServer(http.Dir("./static"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", fs))
	//s.engine.Use(static.Serve("/static/", static.LocalFile("./static", false)))

}

//func (s *Server) registerRoutes() {
//	s.engine.GET("/", s.hh.Home_g())
//	s.engine.GET("/ws", handlers.WsEndpoint())
//
//	//fs := http.FileServer(http.Dir("./static"))
//	//http.Handle("/static/", http.StripPrefix("/static/", fs))
//	s.engine.Use(static.Serve("/static/", static.LocalFile("./static", false)))
//
//}
