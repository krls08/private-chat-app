package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/krls08/private-chat-app/internal/server/handlers"
)

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
