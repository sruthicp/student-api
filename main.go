package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"student-api/config"
	"student-api/controller"
	"student-api/db"
	"student-api/handlers"
	"student-api/repositories"
	"student-api/service"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	pb "student-api/proto/student"
)

func main() {
	// runRestService()
	rungRPCService()
}

func rungRPCService() {
	svc := grpc.NewServer()

	// initializing service config
	config.NewServiceConfig()

	// setting up DB connection
	connection, err := db.NewDBConnection(config.SvcConf)
	if err != nil {
		log.Fatal("postgres connection failure.!!", err)
	}

	sr := repositories.NewStudentRepo(connection)
	ss := service.NewStudentService(sr)
	sc := controller.NewStudentController(ss)

	pb.RegisterStudentServer(svc, sc)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("error listening on port", err)
	}

	svc.Serve(listen)

}

func runRestService() {
	shandler := handlers.NewStudent()

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
