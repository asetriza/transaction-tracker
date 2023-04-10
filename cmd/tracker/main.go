package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/asetriza/transaction-tracker/internal/tracker/adapters/http/rest"
	"github.com/asetriza/transaction-tracker/internal/tracker/app"
	models "github.com/asetriza/transaction-tracker/internal/tracker/models"
	"github.com/asetriza/transaction-tracker/internal/tracker/repository/postgresql"
	"github.com/asetriza/transaction-tracker/internal/tracker/usecase/tracker"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		lg.Info("Initializing",
			zap.String("http.addr", os.Getenv("PORT")),
		)

		dbConfig := postgresql.NewDBConfig()
		repo, err := postgresql.NewClient(dbConfig)
		if err != nil {
			return err
		}
		service := tracker.NewTracker(repo)
		handler := rest.NewHandler(service)
		oasServer, err := models.NewServer(handler)
		if err != nil {
			return err
		}

		srv := rest.NewServer(os.Getenv("PORT"), oasServer)

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			<-ctx.Done()
			return srv.Shutdown(ctx)
		})
		g.Go(func() error {
			defer lg.Info("Server stopped")
			if err := srv.Run(); err != nil && err != http.ErrServerClosed {
				return fmt.Errorf("http %w", err)
			}
			return nil
		})

		return g.Wait()
	})
}
