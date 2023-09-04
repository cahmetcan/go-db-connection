package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

type Row struct {
	id         any
	created_at time.Time
	title      string
	db_time    time.Time
}

func dbConnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseUrl := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	db = conn
	log.Println("Connected to database")

}
