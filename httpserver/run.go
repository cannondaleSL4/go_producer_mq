package httpserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	. "go_producer_mq/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run will run the HTTP HttpServer
func Run(config Config) {
	// Set up a channel to listen to for interrupt signals
	var runChanH = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.HttpServer.Timeout.Server,
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         config.HttpServer.Host + ":" + config.HttpServer.Port,
		Handler:      NewRouter(),
		ReadTimeout:  config.HttpServer.Timeout.Read * time.Second,
		WriteTimeout: config.HttpServer.Timeout.Write * time.Second,
		IdleTimeout:  config.HttpServer.Timeout.Idle * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChanH, os.Interrupt, syscall.SIGTSTP)

	// Alert the user that the server is starting
	log.Printf("HttpServer is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("HttpServer failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChanH

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Printf("HttpServer is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HttpServer was unable to gracefully shutdown due to err: %+v", err)
	}
}

func NewRouter() *http.ServeMux {
	// Create router and define routes and return that router
	router := http.NewServeMux()
	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	return router
}

func RunHttpServer(config Config) {
	// router
	r := mux.NewRouter()
	r.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	log.Printf("INFO: init http api")

	// start server
	err := http.ListenAndServe(":"+config.HttpServer.Port, r)
	if err != nil {
		log.Printf("ERROR: fail init http server, %s", err.Error)
		os.Exit(1)
	}
}
