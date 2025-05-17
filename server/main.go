package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "grpcdemo/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type MessageServer struct {
	pb.UnimplementedMessageServiceServer
	mu       sync.RWMutex
	messages map[string]*pb.GetMessageResponse
}

func NewMessageServer() *MessageServer {
	return &MessageServer{
		messages: make(map[string]*pb.GetMessageResponse),
	}
}

func (s *MessageServer) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	log.Printf("Received GetMessage request for name: %s", req.Name)

	s.mu.RLock()
	message, exists := s.messages[req.Name]
	s.mu.RUnlock()

	if !exists {
		return nil, status.Error(codes.NotFound, "Message not found for name: "+req.Name)
	}

	return message, nil
}

func (s *MessageServer) SubmitMessage(ctx context.Context, req *pb.SubmitMessageRequest) (*pb.SubmitMessageResponse, error) {
	log.Printf("Received SubmitMessage request for name: %s", req.Name)
	fmt.Printf("Received SubmitMessage request for name: %s", req.Name)

	if req.Name == "" || req.Message == "" {
		return nil, status.Error(codes.InvalidArgument, "Name and message are required")
	}

	message := &pb.GetMessageResponse{
		Name:    req.Name,
		Message: req.Message,
	}

	s.mu.Lock()
	s.messages[req.Name] = message
	s.mu.Unlock()

	return &pb.SubmitMessageResponse{
		Success: true,
		Message: "Message submitted successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register our service
	pb.RegisterMessageServiceServer(server, NewMessageServer())

	// Register reflection service on gRPC server
	reflection.Register(server)

	log.Printf("gRPC server starting on :8080...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
