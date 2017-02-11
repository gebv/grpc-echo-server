package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"apps/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	serverSignal := make(chan struct{}, 1)

	// rpc server
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "48655"
		}
		addr := ":" + port
		log.Printf("setup rpc server, address=%s\n", addr)

		tcp, err := net.Listen("tcp", addr)
		if err != nil {
			log.Printf("error listen address=%s. %s\n", addr, err)
			serverSignal <- struct{}{}
			return
		}
		s := grpc.NewServer()
		echo.Setup(s)
		reflection.Register(s)
		if err := s.Serve(tcp); err != nil {
			log.Printf("error start server, address=%s\n", addr)
			serverSignal <- struct{}{}
			return
		}
	}()

	osSignal := make(chan os.Signal, 2)
	close := make(chan struct{})
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	go func() {

		defer func() {
			close <- struct{}{}
		}()

		select {
		case <-osSignal:
			log.Println("signal completion of the process: OS")
		case <-serverSignal:
			log.Println("signal completion of the process: internal (http server, etc..)")
		}

		// TODO: destroy
	}()

	<-close
	os.Exit(0)
}
