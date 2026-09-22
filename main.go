package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// GetMessage returns a greeting message.
func GetMessage(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// HelloHandler handles HTTP requests and responds with the greeting.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Nishchal"
	}
	fmt.Fprintln(w, GetMessage(name))
}

func main() {
	http.HandleFunc("/", HelloHandler)

	// Render dynamically sets the PORT environment variable (e.g. 10000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
