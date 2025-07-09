package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, "Backend is healthy")
	})

	log.Println("Backend is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
