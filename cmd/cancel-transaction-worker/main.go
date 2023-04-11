package main

import (
	"log"
	"time"

	"github.com/asetriza/transaction-tracker/internal/repository/postgresql"
	"github.com/asetriza/transaction-tracker/internal/usecase/account"
	"github.com/asetriza/transaction-tracker/internal/usecase/transaction"
	"github.com/go-co-op/gocron"
)

func main() {
	dbConfig := postgresql.NewDBConfig()
	repo, err := postgresql.NewClient(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s := gocron.NewScheduler(time.UTC)
	account := account.NewAccount(repo)
	transaction := transaction.NewTransactionService(repo, account)

	s.Every(1).Minute().Do(func() {
		err := transaction.CancelLatestOddRecords(10)
		if err != nil {
			log.Printf("error canceling latest odd records %s", err)
		}
	})

	s.StartBlocking()
	log.Println("cron worker job started")
}
