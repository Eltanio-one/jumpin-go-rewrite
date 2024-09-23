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
	"github.com/Eltanio-one/jumpin-go-rewrite/src/session"

	"github.com/joho/godotenv"
)

func main() {
	logger := log.New(os.Stdout, "jumpin-rewrite", log.LstdFlags)

	sessionKey, err := session.GenerateSecureToken(32)
	if err != nil {
		logger.Println("error generating secure session token: ", err)
		return
	}
	os.Setenv("SESSION_KEY", sessionKey)

	err = godotenv.Load()
	if err != nil {
		logger.Println("error loading .env file: ", err)
		return
	}

	store := session.GenerateStore()

	handler.Store = store

	mux := http.NewServeMux()

	mux.Handle("/ping", http.HandlerFunc(handler.Ping))
	mux.Handle("/login", http.HandlerFunc(handler.Login))
	mux.Handle("/register", http.HandlerFunc(handler.Register))
	mux.Handle("/registergym", http.HandlerFunc(handler.RegisterGym))
	mux.Handle("/assigngym", http.HandlerFunc(handler.AssignGym))
	mux.Handle("/sessionplan", http.HandlerFunc(handler.SessionPlan))

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
			logger.Println("error serving due to:", err)
			return
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-signalChannel
	logger.Println("server terminating, gracefully exiting", sig)

	timeoutContext, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() {
		cancelCtx()
	}()

	if err := srvr.Shutdown(timeoutContext); err != nil {
		logger.Fatalf("server shutdown failed:%+s", err)
	}
}
