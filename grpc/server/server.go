package main

import (
	"btime"
	"encoding/json"
	"log"
	"net"

	pb "btime/grpc/btimegrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50055"
)

type resultS struct {
	Ret [5]uint64 `json:"ret"`
	Err error     `json:"err"`
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetBinary(ctx context.Context, in *pb.TimeRequest) (*pb.BinaryReply, error) {
	data, _ := btime.GetBinary(in.StartTime, in.EndTime)
	result := resultS{data, nil}
	json, _ := json.Marshal(result)
	str := "startTime: " + in.StartTime + ",endtime: " + in.EndTime + "json: " + string(json)
	return &pb.BinaryReply{JsonMessage: str}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBtimeServer(s, &server{})
	s.Serve(lis)
}
