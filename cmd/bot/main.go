package main

import (
	"route256/config"
	"route256/internal/clients/tg"
	"route256/internal/model/messages"
	"route256/internal/model/spending"
	"route256/internal/storage/spending_storage"
)

func main() {
	config, err := config.New()
	if err != nil {
		//logger_custome.Logger.Fatal(err)
	}

	tgClient, err := tg.Start(config)
	if err != nil {
		//logger_custome.Logger.Fatal(err)
	}

	spendingStore := spending_storage.New()
	spendingModel := spending.New(spendingStore)

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel, spendingModel)
}
