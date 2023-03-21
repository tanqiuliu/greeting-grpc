package server

import (
	"context"
	"fmt"
	pb "github.com/tanqiuliu/greeting-grpc/api/greetingpb"
)

const greetingMessageTemplate = "Hello, %s"

type GreetingServerImpl struct {
	pb.UnimplementedGreetingServer
}

func (server *GreetingServerImpl) SayHello(_ context.Context, req *pb.GreetingRequest) (*pb.GreetingReply, error) {
	var message = fmt.Sprintf(greetingMessageTemplate, req.Name)
	return &pb.GreetingReply{GreetingMessage: message}, nil
}

func NewGreetingServerImpl() pb.GreetingServer {
	return &GreetingServerImpl{}
}
