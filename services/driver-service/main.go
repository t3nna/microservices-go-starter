package main

import (
	"context"
	grpcserver "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"ride-sharing/shared/env"
	"ride-sharing/shared/messaging"
	"syscall"
)

var GrpcAddr = ":9092"

func main() {
	rabbitMqUri := env.GetString("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := NewService()
	// rabbitMQ connection
	rabbitmq, err := messaging.NewRabbitMQ(rabbitMqUri)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close()

	// starting the grpcServer
	grpcServer := grpcserver.NewServer()
	NewGrpcHandler(grpcServer, svc)

	log.Println("Starting gRPC server Driver service on port ", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()

	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()

}
