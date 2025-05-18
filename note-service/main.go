package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	userclient "note-service/grpc" 
)

var client *userclient.UserClient

func main() {
	
	time.Sleep(3 * time.Second) 

	conn, err := grpc.Dial("user-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Failed to connect to user-service: %v", err)
	}
	defer conn.Close()

	client = userclient.NewUserClient(conn)


	http.HandleFunc("/notes/", handleNote)
	log.Println("📡 NoteService REST API is running on port 50051...")
	log.Fatal(http.ListenAndServe(":50051", nil))
}

func handleNote(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/notes/"):] 


	res, err := client.GetUser(id)
	if err != nil {
		http.Error(w, "❌ failed to get user", 500)
		return
	}

	
	json.NewEncoder(w).Encode(res)
}
