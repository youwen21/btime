package main

import (
	"log"
	"os"

	pb "btime/grpc/btimegrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address          = "localhost:50055"
	defaultStartTime = "08:00:00"
	defaultEndTime   = "09:00:00"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBtimeClient(conn)

	// Contact the server and print out its response.
	startTime := defaultStartTime
	if len(os.Args) > 1 {
		startTime = os.Args[1]
	}
	endTime := defaultEndTime
	if len(os.Args) > 2 {
		endTime = os.Args[2]
	}
	r, err := c.GetBinary(context.Background(), &pb.TimeRequest{StartTime: startTime, EndTime: endTime})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.JsonMessage)
}
