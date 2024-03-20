package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"grpc.worklen.com/server"

	"github.com/pkg/errors"
)

func runServer(log *log.Logger) error {
	log.Println("main: Initializing GRPC server")
	defer log.Println("main: Completed")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	serverErrors := make(chan error, 1)
	server, err := server.StartServer()
	if err != nil {
		return errors.Wrap(err, "running server")
	}

	go func() {
		log.Printf("main: GRPC server listening")
		serverErrors <- server.Serve()
	}()

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main: %v: Start shutdown", sig)
		server.GracefulStop()
	}

	return nil
}

func main() {
	log := log.New(os.Stdout, "GRPC SERVER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := runServer(log); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
