package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Skills    []string  `json:"skills"` // <--- NEW
	CreatedAt time.Time `json:"created_at"`
}
