package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

type Row struct {
	id         any
	created_at time.Time
	title      string
	db_time    time.Time
}

func dbConnect() (*pgxpool.Pool, *pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	databaseUrl := os.Getenv("DATABASE_URL")

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, nil, err
	}
	fmt.Println("Connected to database")

	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to acquire connection from pool: %v\n", err)
		return nil, nil, err
	}

	return dbPool, conn.Conn(), nil
}
