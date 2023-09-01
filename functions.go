package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func execQuery(query string) ([]interface{}, string, string) {

	connTime := time.Now()
	_, conn, err := dbConnect()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic(err)
	}

	startQuery := time.Now()
	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err.Error(), "ERROR"
	}

	values := make([]interface{}, 0)
	for rows.Next() {
		value, err := rows.Values()
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return nil, err.Error(), "ERROR"
		}
		values = append(values, value...)
	}

	return values, (time.Since(startQuery)).String(), (time.Since(connTime) - time.Since(startQuery)).String()
}

// must be removed
func count(query string, Conn *pgx.Conn) (any, int, string) {

	startQuery := time.Now()
	rows, err := Conn.Query(context.Background(), query)
	execTime := time.Since(startQuery)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, 0, "ERROR"
	}
	defer rows.Close()

	var rowSlice []Row
	for rows.Next() {
		var r Row
		err := rows.Scan(&r.id, &r.created_at, &r.title, &r.db_time)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error scanning rows:", err)
	}

	return rows, 1, execTime.String()
}

func poolQuery(query string, dbPool *pgxpool.Pool) (int64, string) {
	startQuery := time.Now()
	rows, err := dbPool.Exec(context.Background(), query)
	execTime := time.Since(startQuery)

	if err != nil {
		fmt.Println("Error executing query:!", err)
		return 0, "0"
	}

	return rows.RowsAffected(), execTime.String()
}
