package main

import (
	"fmt"
	"os"

	// "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong ")
	})
	/* 	r.GET("/qE", query)
	   	r.GET("/insert", insertRandomId)
	   	r.GET("/getCount", getCount) */

	// ÖZGÜR BEY'İN YAZDIKLARI
	r.GET("/execute", executeParam) // execute?query=select * from test
	r.GET("/getCountsByParsing", getCountsByParsing)
	r.GET("/getRowsWithLimit", getRowsWithLimit) // getRowsWithLimit?limit=10
	r.GET("/getMaxId", getMaxId)
	// r.GET("/getRandom", getRandomRowById)

	go func() {
		if err := r.Run(":9000"); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		}
		fmt.Println("Server started")
	}()
	select {}
}
