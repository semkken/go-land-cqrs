package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	commandHandler := commands.NewCommandHandler()
	queryHandler := queries.NewQueryHandler()

	server := &http.Server{
		Addr:         ":8080",                                 // Server address
		Handler:      newRouter(commandHandler, queryHandler), // Custom router for handling requests
		ReadTimeout:  15 * time.Second,                        // Maximum duration before timing out read operations
		WriteTimeout: 15 * time.Second,                        // Maximum duration before timing out write operations
		IdleTimeout:  60 * time.Second,                        // Maximum duration for which the server will keep connections open
	}

	go func() {
		fmt.Println("Server is running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	waitForTermination(server)
}

func newRouter(commandHandler *commands.CommandHandler, queryHandler *queries.QueryHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/books", commandHandler.AddBookHandler).Methods("POST")
	router.HandleFunc("/api/books/{id}", commandHandler.UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/api/books/{id}", commandHandler.RemoveBookHandler).Methods("DELETE")
	router.HandleFunc("/api/books/{id}", queryHandler.GetBookByIDHandler).Methods("GET")
	router.HandleFunc("/api/books", queryHandler.GetAllBooksHandler).Methods("GET")
	return router
}

func waitForTermination(server *http.Server) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	<-signalCh

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	fmt.Println("Server shut down gracefully")
}
