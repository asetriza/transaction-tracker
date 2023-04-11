package main

import (
	"context"
	"log"
	"time"

	"github.com/asetriza/transaction-tracker/internal/repository/postgresql"
	accountrepo "github.com/asetriza/transaction-tracker/internal/repository/postgresql/account-repo"
	transactionrepo "github.com/asetriza/transaction-tracker/internal/repository/postgresql/transaction-repo"
	"github.com/asetriza/transaction-tracker/internal/usecase/account"
	"github.com/asetriza/transaction-tracker/internal/usecase/transaction"
	"github.com/go-co-op/gocron"
)

func main() {
	config := postgresql.NewConfig()
	client, err := postgresql.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	s := gocron.NewScheduler(time.UTC)
	accountRepo := accountrepo.New(client)
	accountService := account.NewAccountService(accountRepo)
	transactionRepo := transactionrepo.New(client)
	transactionService := transaction.NewTransactionService(transactionRepo, accountService)

	s.Every(1).Minute().Do(func() {
		err := transactionService.CancelLatestOddRecords(context.TODO(), 10)
		if err != nil {
			log.Printf("error canceling latest odd records %s", err)
		}
	})

	s.StartBlocking()
	log.Println("cron worker job started")
}
