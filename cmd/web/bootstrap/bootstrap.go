package bootstrap

import (
	"log"
	"net/http"

	"github.com/krls08/private-chat-app/cmd/web/routes"
)

func Run() error {
	mux := routes.Routes()

	log.Println("Starting webserver on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}
	return nil
}
