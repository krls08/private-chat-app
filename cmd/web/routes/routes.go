package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/krls08/private-chat-app/internal/infrastructure/server/handlers"
)

func Mux(h handlers.HomeHandlers) http.Handler {
	mux := pat.New()

	//mux.Get("/", http.HandlerFunc(handlers.Home))

	//mux.Get("/", http.HandlerFunc(h.Home))
	//mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
