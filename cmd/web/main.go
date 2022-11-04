package main

import (
	"log"

	"github.com/krls08/private-chat-app/cmd/web/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
