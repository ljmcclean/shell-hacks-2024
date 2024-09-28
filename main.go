package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const port = ":3000"
const killTime = 10 // time allotted for shutdown in seconds

func main() {
	ctx := context.Background()

	err := run(ctx)
	if err != nil {
		log.Fatalf("server closed with error: %s\n", err)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	server := NewServer(port)

	go func() {
		log.Printf("listening and serving on %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, killTime*time.Second)
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}
