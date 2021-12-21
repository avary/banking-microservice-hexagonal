package main

import (
	"context"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/handlers"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const port = ":5000"

func main() {
	l := log.New(os.Stdout, "banking-server ", log.LstdFlags)

	// wire up the handlers
	dbClient := service.GetDbClient(l)
	customerDbConn := domain.NewCustomerRepoDb(dbClient, l)
	accountDbConn := domain.NewAccountRepoDb(dbClient, l)
	ch := handlers.CustomerHandlers{Service: service.NewCustomerService(customerDbConn), L: l}
	ah := handlers.AccountHandlers{Service: service.NewAccountService(accountDbConn), L: l}

	// create a router and register handlers
	r := mux.NewRouter()
	getRtr := r.Methods(http.MethodGet).Subrouter()
	getRtr.HandleFunc("/customers/", ch.GetAllCustomers)
	getRtr.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerByID)

	postRtr := r.Methods(http.MethodPost).Subrouter()
	postRtr.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount)

	// creating the server
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	// go routine to start server on port 8080
	go startServer(srv, l)

	// wait for interrupt signal to gracefully shut down the server with a timeout of 30 seconds.
	gracefulShutdown(srv, l)
}

// start server
func startServer(srv *http.Server, l *log.Logger) {
	l.Printf("Starting server on port %s", port)
	if err := srv.ListenAndServe(); err != nil {
		l.Printf("Error starting server: %s", err)
	}
}

// wait for interrupt signal to gracefully shut down the server with a timeout of 30 seconds.
func gracefulShutdown(srv *http.Server, l *log.Logger) {
	// listen for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// graceful shutdown
	l.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		l.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	l.Println("Server gracefully stopped")
}
