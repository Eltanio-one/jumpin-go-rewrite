package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Eltanio-one/jumpin-go-rewrite/src/handler"
)

func main() {
	logger := log.New(os.Stdout, "jumpin-rewrite", log.LstdFlags)

	mux := http.NewServeMux()

	mux.Handle("/ping", http.HandlerFunc(handler.Ping))
	mux.Handle("/login", http.HandlerFunc(handler.Login))

	srvr := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := srvr.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Println("Error serving due to:", err)
			return
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-signalChannel
	logger.Println("Server terminating, gracefully exiting", sig)

	timeoutContext, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() {
		cancelCtx()
	}()

	if err := srvr.Shutdown(timeoutContext); err != nil {
		logger.Fatalf("server shutdown failed:%+s", err)
	}
}
