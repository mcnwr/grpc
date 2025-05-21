package main

import (
	"context"
	"log"
	"time"

	pb "grpcdemo/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewMessageServiceClient(conn)

	// First, submit a message
	submitResp, err := client.SubmitMessage(ctx, &pb.SubmitMessageRequest{
		Id:      "N250",
		Message: "N250 Gatotkaca adalah pesawat terbang pertama buatan Indonesia. Pesawat ini dirancang dan diprakarsai oleh B.J. Habibie, seorang ilmuan dirgantara yang diakui di dunia. N250 Gatotkaca diproduksi oleh Industri Pesawat Terbang Nusantara (IPTN), yang kemudian berubah nama menjadi PT Dirgantara Indonesia.",
	})
	if err != nil {
		log.Fatalf("could not submit message: %v", err)
	}
	log.Printf("Submit Message Response: %v", submitResp)

	// Then, get the message
	getResp, err := client.GetMessage(ctx, &pb.GetMessageRequest{
		Id: "N250",
	})
	if err != nil {
		log.Fatalf("could not get message: %v", err)
	}
	log.Printf("Get Message Response: %v", getResp)
}
