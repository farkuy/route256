package main

import (
	"log"
	"route256/config"
	"route256/internal/clients/tg"
	"route256/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	tgClient, err := tg.Start(config)
	if err != nil {
		log.Fatal(err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
