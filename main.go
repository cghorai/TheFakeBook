package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	grpcserver "projects/TheFakeBook/internal/server"
)

const (
	grpcPort           = 9981
	gatewayServicePort = 8083
	host               = "localhost"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup mongo connection
	mongoConnectionString := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionString))
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB %v, error: %v", mongoConnectionString, err)
	}

	// start grpc server
	log.Printf("GRPC Server listenning on port %v", grpcPort)
	go grpcserver.StartService(grpcserver.StartServiceInput{
		GrpcPort:        grpcPort,
		RestServicePort: gatewayServicePort,
		MongoClient:     client,
	})

	//start REST Gateway server
	log.Printf("REST API server listenning on port %v", gatewayServicePort)
	grpcserver.StartGatewayProxy(grpcserver.GrpcGatewayParams{
		ServicePort: gatewayServicePort,
		GrpcPort:    grpcPort,
		Host:        host})

}
