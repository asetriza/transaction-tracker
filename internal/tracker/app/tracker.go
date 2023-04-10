package app

import (
	"context"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

// func Run() error {
// 	dbConfig := postgresql.NewDBConfig()
// 	repo, err := postgresql.NewClient(dbConfig)
// 	if err != nil {
// 		return err
// 	}
// 	service := tracker.NewTracker(repo)
// 	handler := rest.NewHandler(service)
// 	srv, err := models.NewServer(handler)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := http.ListenAndServe(os.Getenv("PORT"), srv); err != nil {
// 		return err
// 	}

// 	srvs := rest.NewServer(os.Getenv("PORT"), srv)

// 	go func() {
// 		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
// 			log.Println(err)
// 		}
// 	}()

// 	log.Println("Server started")

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
// 	<-quit

// 	const timeout = 5 * time.Second
// 	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
// 	defer shutdown()

// 	if err := srv.(ctx); err != nil {
// 		return fmt.Errorf("failed to stop server: %v", err)
// 	}

// 	if err := conn.Close(); err != nil {
// 		return err
// 	}

// 	return nil

// 	return nil
// }

const EnvLogLevel = "LOG_LEVEL"

const (
	exitCodeOk       = 0
	exitCodeWatchdog = 1
)

const (
	shutdownTimeout = time.Second * 5
	watchdogTimeout = shutdownTimeout + time.Second*5
)

// Run f until interrupt.
func Run(f func(ctx context.Context, log *zap.Logger) error) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := zap.NewProductionConfig()

	if s := os.Getenv(EnvLogLevel); s != "" {
		var lvl zapcore.Level
		if err := lvl.UnmarshalText([]byte(s)); err != nil {
			panic(err)
		}
		cfg.Level.SetLevel(lvl)
	}

	lg, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		if err := f(ctx, lg); err != nil {
			return err
		}
		return nil
	})
	go func() {
		// Guaranteed way to kill application.
		<-ctx.Done()

		// Context is canceled, giving application time to shut down gracefully.
		lg.Info("Waiting for application shutdown")
		time.Sleep(watchdogTimeout)

		// Probably deadlock, forcing shutdown.
		lg.Warn("Graceful shutdown watchdog triggered: forcing shutdown")
		os.Exit(exitCodeWatchdog)
	}()

	if err := wg.Wait(); err != nil {
		if err == context.Canceled {
			lg.Info("Graceful shutdown")
			return
		}
		lg.Fatal("Failed",
			zap.Error(err),
		)
	}

	os.Exit(exitCodeOk)
}
