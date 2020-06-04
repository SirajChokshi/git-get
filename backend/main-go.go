package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
)

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "git-get"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("not created!")
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

type Data struct {
	User *User `json:"user"`
}

type User struct {
	Object string `json:"data"`
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Make the app context and client
	ctx := context.Background()
	client := createClient(ctx)
	// Close client on application end
	defer client.Close()
	// Add data
	//_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
	//	"first": "Ada",
	//	"last":  "Lovelace",
	//	"born":  1815,
	//})
	//// If there's any error, log it.
	//if err != nil {
	//	log.Fatalf("Failed adding alovelace: %v", err)
	//}

	url := r.RequestURI
	// Get the name from the string
	params := strings.Split(url, "/")
	userRef := client.Collection("users")
	if params[2] != "" {
		// Get user from firestore
		doc, _ := userRef.Doc(params[2]).Get(ctx)
		if doc.Data() == nil {
			// No user exists in firebase, make new one.
			println("Creating user %s", params[2])
			jsonUser := makeRequest(params[2])
			// The Set() command either creates a user or updates the user.
			_, err := userRef.Doc(params[2]).Set(ctx, jsonUser)
			// Handle the error
			if err != nil {
				log.Fatalf("Failed to add user: %s", params[2])
			} else {
				json.NewEncoder(w).Encode(jsonUser)
			}
		} else {
			// Check the last updated timestamp
			lastUpdated := doc.UpdateTime
			if lastUpdated.Sub(time.Now()).Hours() <= 6 {
				json.NewEncoder(w).Encode(doc.Data())
			} else {
				// Make the request if they haven't been updated in 6 hours
				println("Updating user %s, last update: %s", params[2], lastUpdated.String())
				jsonUser := makeRequest(params[2])
				_, err := userRef.Doc(params[2]).Set(ctx, jsonUser)
				if err != nil {
					log.Fatalf("Failed to add user: %s", params[2])
				} else {
					json.NewEncoder(w).Encode(map[string]interface{}{"response": jsonUser})
				}
			}
		}
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"user": nil})
	}

}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get/", get)
	http.HandleFunc("/ratelimit/", getRemainingRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
