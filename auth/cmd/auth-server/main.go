package main

import (
	"context"
	"github.com/ashtishad/banking-microservice-hexagonal/auth/internal/handlers"
	"github.com/ashtishad/banking-microservice-hexagonal/auth/pkg/domain"
	"github.com/ashtishad/banking-microservice-hexagonal/auth/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "auth-server ", log.LstdFlags)
	router := mux.NewRouter()
	authRepository := domain.NewAuthRepository(service.GetDbClient())
	authHandler := handlers.AuthHandler{Service: service.NewLoginService(authRepository, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", authHandler.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", authHandler.Verify).Methods(http.MethodGet)

	// creating the server
	srv := &http.Server{
		Addr:         ":5001",
		Handler:      router,
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
	l.Printf("Starting server on port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		l.Printf("%s", err.Error())
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
