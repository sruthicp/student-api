package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"student-api/db"
	"student-api/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var postgres *db.DBConnection
	var err error

	postgres, err = db.NewDBConnection()

	if err != nil {
		panic(err)
	} else if postgres == nil {
		panic("postgres is nil")
	}

	shandler := handlers.NewStudent(postgres)

	sm := mux.NewRouter()
	sm.Handle("/students", shandler).Methods("POST")
	sm.Handle("/students/{id}", shandler).Methods("GET", "PUT", "DELETE")

	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	out := <-sigChan
	log.Println("graceful shutdown of service", out)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	cancel()
	s.Shutdown(ctx)
}
