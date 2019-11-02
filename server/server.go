package main

import (
	"io"
	"log"
	"net"

	greetpb "github.com/Rajat2019/GRPC_IN_ACTION/04-BiDirectionalStreaming/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(stream greetpb.GreeteveryOneService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sendErr := stream.Send(&greetpb.GreetEveryOneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("this is an error %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreeteveryOneServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
