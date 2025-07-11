package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/omar212/SkillSwap/backend/db"
	"github.com/omar212/SkillSwap/backend/handlers"
	"github.com/omar212/SkillSwap/backend/middleware"
)

func main() {
	// Load env and connect DB
	_ = godotenv.Load()
	db.ConnectDB()
	db.CreateUsersTable()
	db.SeedUsers()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend connected and users seeded 🚀"))
	})

	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetUsersHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateUserHandler(w, r)
		} else if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Wrap mux with CORS middleware
	handler := middleware.CORSMiddleware(mux)

	log.Println("✅ Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
