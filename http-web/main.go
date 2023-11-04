package main

import (
	"context"
	"http-web/server"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// domains := []string{"example.com", "www.example.com"}
	handler := server.NewRouter()

	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
		//	ReadTimeout:  5 * time.Second,
		//	WriteTimeout: 10 * time.Second,
		//	IdleTimeout:  120 * time.Second,
	}

	go func() {
		srv.ListenAndServe()
	}()

	с := make(chan os.Signal, 1)
	signal.Notify(с, os.Interrupt)
	<-с

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
