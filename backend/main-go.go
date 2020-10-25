package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "Hello, World!")
}

type Data struct {
	User *User `json:"user"`
}

type User struct {
	Object string `json:"data"`
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Make the app context and client
	ctx := context.Background()
	client := createClient(ctx)
	// Close client on application end
	defer client.Close()
	// Add data

	url := r.RequestURI
	// Get the name from the string
	params := strings.Split(url, "/")
	userRef := client.Collection("users")
	if params[2] != "" {
		// Get user from firestore
		lowercase := strings.ToLower(params[2])
		doc, _ := userRef.Doc(lowercase).Get(ctx)
		if doc.Data() == nil {
			// No user exists in firebase, make new 5000one.
			println("Creating user %s", lowercase)
			jsonUser := makeRequest(lowercase)
			// The Set() command either creates a user or updates the user.
			_, err := userRef.Doc(lowercase).Set(ctx, jsonUser)
			println("")
			// Handle the error
			if err != nil {
				log.Fatalf("Failed to add user: %s", lowercase)
			} else {
				json.NewEncoder(w).Encode(jsonUser)
			}
		} else {
			// Make the request if they haven't been updated in 6 hours
			println("Updating user %s", lowercase)
			jsonUser := makeRequest(lowercase)
			_, err := userRef.Doc(lowercase).Set(ctx, jsonUser)
			if err != nil {
				log.Fatalf("Failed to add user: %s", lowercase)
			} else {
				json.NewEncoder(w).Encode(map[string]interface{}{"response": jsonUser})
			}
		}
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"user": nil})
	}

}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
		lowercase := strings.ToLower(params[2])
		doc, _ := userRef.Doc(lowercase).Get(ctx)
		if doc.Data() == nil {
			// No user exists in firebase, make new 5000one.
			println("Creating user %s", lowercase)
			jsonUser := makeRequest(lowercase)
			// The Set() command either creates a user or updates the user.
			_, err := userRef.Doc(lowercase).Set(ctx, jsonUser)
			println("")
			// Handle the error
			if err != nil {
				log.Fatalf("Failed to add user: %s", lowercase)
			} else {
				json.NewEncoder(w).Encode(jsonUser)
			}
		} else {
			// Check the last updated timestamp
			lastUpdated := doc.UpdateTime
			if time.Now().Sub(lastUpdated).Hours() <= 6 {
				println(lastUpdated.Sub(time.Now()).Hours())
				json.NewEncoder(w).Encode(doc.Data())
			} else {
				// Make the request if they haven't been updated in 6 hours
				println("Updating user %s, last update: %s", lowercase, lastUpdated.String())
				jsonUser := makeRequest(lowercase)
				_, err := userRef.Doc(lowercase).Set(ctx, jsonUser)
				if err != nil {
					log.Fatalf("Failed to add user: %s", lowercase)
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
	http.HandleFunc("/update/", update)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	fmt.Println(port)
	return ":" + port
}

func main() {
	handleRequest()
}
