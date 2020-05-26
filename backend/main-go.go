package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/iterator"

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

func test(w http.ResponseWriter, r *http.Request) {
	// Make the app context and client
	ctx := context.Background()
	client := createClient(ctx)
	// Close client on application end
	defer client.Close()

	// Add data
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	// If there's any error, log it.
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	// Get all users from firestore
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
