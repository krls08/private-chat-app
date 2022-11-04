package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/krls08/private-chat-app/internal/server/handlers"
)

func Mux() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
