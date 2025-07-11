package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omar212/SkillSwap/backend/models"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	Pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		panic("Unable to connect to DB: " + err.Error())
	}

	err = Pool.Ping(ctx)
	if err != nil {
		panic("Database ping failed: " + err.Error())
	}

	fmt.Println("✅ Connected to AWS RDS PostgreSQL")
}

func CreateUsersTable() error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW()
    );`

	ctx := context.Background()

	_, err := Pool.Exec(ctx, createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	log.Println("✅ users table created or already exists")
	return nil
}

func InsertUser(user *models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`

	ctx := context.Background()
	row := Pool.QueryRow(ctx, query, user.Name, user.Email)

	err := row.Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]models.User, error) {
	query := `SELECT id, name, email, created_at FROM users`

	ctx := context.Background()
	rows, err := Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
