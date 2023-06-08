package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	defaultPort = 8080
)

func parseEnv() int64 {
	env := os.Getenv("PORT")
	if env == "" {
		return defaultPort
	}

	port, err := strconv.ParseInt(env, 10, 32)
	if err != nil {
		return defaultPort
	}

	return port
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("received request for url: %s from %s\n",
		req.URL.String(),
		req.RemoteAddr,
	)

	host, err := os.Hostname()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(res, "Hello, World!\nServer: %s\n", host)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRequest))

	port := parseEnv()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM)

	go func() {
		log.Printf("server listening on port %d", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error: %v", err)
		}
	}()

	sig := <-quit
	log.Printf("received signal %d", sig)
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("can't shut down: %v", err)
	}
}
