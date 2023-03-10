package cli

import (
	"context"
	"fmt"
	"github.com/rubiskko/greeting-grpc/api/greetingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func Run(name string) error {
	// init client stub
	conn, err := grpc.Dial("localhost:50001",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	conn.GetState()
	defer conn.Close()
	greetingClient := greetingpb.NewGreetingClient(conn)

	// call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	message, nil := greetingClient.SayHello(ctx, &greetingpb.GreetingRequest{Name: name})
	if err != nil {
		fmt.Printf("Encountered error when calling SayHello: %s", err)
		return err
	}
	fmt.Printf("Received reply: %s", message)
	return nil
}
