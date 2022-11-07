package bootstrap

import (
	"log"
	"net/http"

	"github.com/krls08/private-chat-app/cmd/web/routes"
	"github.com/krls08/private-chat-app/internal/infrastructure/server/handlers"
)

func Run() error {
	mux := routes.Mux()

	go handlers.ListenToWsChannel()

	log.Println("Starting webserver on port 60002")

	err := http.ListenAndServe(":60002", mux)
	if err != nil {
		return err
	}
	return nil
}
