package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/BigDwarf/testci/internal/application"

	// #nosec G108
	_ "net/http/pprof"
)

func main() {
	ctx := context.Background()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	app := application.NewApp()
	if err := app.Start(); err != nil {
		panic("unimpemented")
	}
	<-done

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	app.Stop(ctx)
}
