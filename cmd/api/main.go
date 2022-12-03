package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/di"
)

func main() {
	app, cleanup, err := di.InitializeApp()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			app.Log.Error("panic", zap.Any("details", r))
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		app.Log.Info("Received a signal.", zap.String("signal", x.String()))
	case e := <-app.Server.Notify():
		app.Log.Error("Received an error from server.", zap.Error(e))
	}

	cleanup()
}
