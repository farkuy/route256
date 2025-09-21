package main

import (
	"log"
	"route256/config"
	"route256/internal/clients/tg"
	"route256/internal/model/messages"
	"route256/internal/model/spending"
	"route256/internal/storage/spending_storage"
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

	spendingStore := spending_storage.New()
	spendingModel := spending.New(spendingStore)

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel, spendingModel)
}
