package main

import (
	"fmt"
	pb "github.com/anupcowkur/grpc_demo_go/timer"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

type server struct {
	time []*pb.TimeRequest
}

// We should be configuring keep alive options so that we don't stream to dead 
// clients but it's a demo, so whatevs.
func (s *server) Timer(request *pb.TimeRequest, stream pb.Timer_TimerServer) error {
	for range time.Tick(time.Second * 1) {
		timeResponse := pb.TimeResponse{Time: time.Now().Format(time.RFC850)}
		if err := stream.Send(&timeResponse); err != nil {
			return err
		}
		fmt.Println("sent:", timeResponse.Time)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterTimerServer(s, &server{})
	s.Serve(lis)
}
