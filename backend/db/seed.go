package db

import (
	"context"
	"fmt"
	"time"
)

func SeedUsers() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Example users to insert
	users := []struct {
		Name   string
		Email  string
		Skills []string
	}{
		{"Alice Johnson", "alice@example.com", []string{"Go", "React"}},
		{"Bob Smith", "bob@example.com", []string{"TypeScript", "Docker"}},
		{"Carol Lee", "carol@example.com", []string{"PostgreSQL", "AWS"}},
	}

	for _, u := range users {
		skills := fmt.Sprintf("{%s}", joinStrings(u.Skills, ","))
		_, err := Pool.Exec(ctx, "INSERT INTO users (name, email, skills) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", u.Name, u.Email, skills)
		if err != nil {
			return fmt.Errorf("failed to insert user %s: %w", u.Email, err)
		}
	}

	fmt.Println("✅ Seeded users table with initial data")
	return nil
}

// Helper function to join string slice with quotes around each skill for PostgreSQL array
func joinStrings(items []string, sep string) string {
	out := ""
	for i, s := range items {
		if i > 0 {
			out += sep
		}
		out += "\"" + s + "\""
	}
	return out
}
