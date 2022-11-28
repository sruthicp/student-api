package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"student-api/config"
	"student-api/controller"
	"student-api/db"
	"student-api/repositories"
	"student-api/service"

	pb "student-api/proto/student"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
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

	go runRestService(sc)
	rungRPCService(sc)
}

func rungRPCService(sc pb.StudentServer) {
	svc := grpc.NewServer()

	pb.RegisterStudentServer(svc, sc)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("error listening on port", err)
	}

	svc.Serve(listen)

}

func runRestService(sc pb.StudentServer) {
	mux := runtime.NewServeMux()

	pb.RegisterStudentHandlerServer(context.Background(), mux, sc)

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}

	// shandler := handlers.NewStudent()

	// sm := mux.NewRouter()
	// sm.Handle("/students", shandler).Methods("POST")
	// sm.Handle("/students/{id}", shandler).Methods("GET", "PUT", "DELETE")

	// s := &http.Server{
	// 	Addr:         "0.0.0.0:8080",
	// 	Handler:      sm,
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  300 * time.Second,
	// 	WriteTimeout: 300 * time.Second,
	// }
	// go func() {
	// 	if err := s.ListenAndServe(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, os.Interrupt)

	// out := <-sigChan
	// log.Println("graceful shutdown of service", out)

	// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// cancel()
	// s.Shutdown(ctx)
}
