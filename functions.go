package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Result struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Db_time    time.Time `json:"db_time"`
}

type Result2 struct {
	Maxid      int64     `json:"maxid"`
	Created_at time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Db_time    time.Time `json:"db_time"`
}

func execQuery(query string) ([]Result, string, string) {

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
		panic(err)
	}

	var rowSlice []Result
	for rows.Next() {
		var r Result
		err := rows.Scan(&r.Id, &r.Created_at, &r.Title, &r.Db_time)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			panic(err)
		}
		rowSlice = append(rowSlice, r)
	}

	return rowSlice, (time.Since(startQuery)).String(), (time.Since(connTime) - time.Since(startQuery)).String()
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

func maxId(tableName string) (Result2, string, string) {
	fmt.Println("maxId", tableName)

	connTime := time.Now()
	_, conn, err := dbConnect()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic(err)
	}

	startQuery := time.Now()
	rows, err := conn.Query(context.Background(), "Select  id maxID, created_at, title, CURRENT_TIMESTAMP  db_time FROM test_table where id = (select max(id) from test_table)")

	if err != nil {
		fmt.Println("Error executing query:", err)
		panic(err)
	}

	var rowSlice Result2
	for rows.Next() {
		var r Result2
		fmt.Println("rows.Next()", rows)
		err := rows.Scan(&r.Maxid, &r.Created_at, &r.Title, &r.Db_time)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			panic(err)
		}
		rowSlice = r
	}

	return rowSlice, (time.Since(startQuery)).String(), (time.Since(connTime) - time.Since(startQuery)).String()

}
