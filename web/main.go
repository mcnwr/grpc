package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	pb "grpcdemo/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PageData struct {
	Name    string
	Message string
}

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := pb.NewMessageServiceClient(conn)

	// Load HTML template
	tmpl := template.Must(template.ParseFiles("../templates/grpc_test.html"))

	// Handle root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		tmpl.Execute(w, PageData{})
	})

	// Handle submit message
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.SubmitMessage(ctx, &pb.SubmitMessageRequest{
			Name:    data.Name,
			Message: data.Message,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Handle get message
	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Path[len("/get/"):]
		if name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.GetMessage(ctx, &pb.GetMessageRequest{
			Name: name,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Start server
	log.Println("Web server starting on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
