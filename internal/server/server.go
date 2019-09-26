package server

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"projects/TheFakeBook/pkg/handler"
	proto "projects/TheFakeBook/pkg/service"
	"strconv"
)

type server struct {
	mongoClient    *mongo.Client
	serviceHandler *handler.FakeBookHandler
}

type StartServiceInput struct {
	MongoClient     *mongo.Client
	GrpcPort        int32
	RestServicePort int32
}

func StartService(in StartServiceInput) {

	s := grpc.NewServer()
	//Injecting the necessary to DB Layer
	handler := &handler.FakeBookHandler{}
	//TODO: We will inject other attributes later as necessary

	proto.RegisterFakeBookServiceServer(s, &server{mongoClient: in.MongoClient, serviceHandler: handler})
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(9981))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) AddFakeNews(ctx context.Context, input *proto.FakeNewsRequestWrapper) (*proto.FakeNewsResponse, error) {
	panic("implement me")
}
func (s *server) DeleteFakeNews(ctx context.Context, input *proto.FakeNewsRequestWrapper) (*proto.FakeNewsResponse, error) {
	panic("implement me")
}
func (s *server) RateFakeNews(ctx context.Context, input *proto.RatingRequestWrapper) (*proto.FakeNewsResponse, error) {
	panic("implement me")
}
func (s *server) ViewFakeNews(ctx context.Context, input *proto.FakeNewsId) (*proto.FakeNewsResponse, error) {
	panic("implement me")
}
func (s *server) HealthCheck(ctx context.Context, input *proto.HealthRequest) (*proto.HealthReply, error) {
	return &proto.HealthReply{Status: "Up"}, nil
}
