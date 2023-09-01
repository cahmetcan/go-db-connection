package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Executes the query given in the parameter
func executeParam(c *gin.Context) {
	tableName := c.Query("q")
	startTime := time.Now()

	rows, execTime, a := execQuery(tableName)

	c.JSON(200, gin.H{
		"success":          "true",
		"endedIn":          time.Since(startTime).String(),
		"queryExecTime":    execTime,
		"databaseConnTime": a,
		"message":          rows,
	})
}

// Gets the counts of the rows in the table
func getCountsByParsing(c *gin.Context) {
	query := "Select  id maxID, created_at, title, CURRENT_TIMESTAMP  db_time FROM test_table where id = (select max(id) from test_table)"

	startTime := time.Now()

	rows, execTime, a := execQuery(query)

	c.JSON(200, gin.H{
		"success":          "true",
		"endPointCallTime": time.Since(startTime).String(),
		"queryExecTime":    execTime,
		"databaseConnTime": a,
		"message":          rows,
	})
}

/*
{
  "success": true,
  "result": [
    {
      "id": 298227258,
      "created_at": "2023-08-28T21:00:00.000Z",
      "title": "sansli numaram 83405",
      "db_time": "2023-09-01T21:55:22.790Z"
    },
    {
      "id": 298227257,
      "created_at": "2023-08-28T21:00:00.000Z",
      "title": "sansli numaram 916151",
      "db_time": "2023-09-01T21:55:22.790Z"
    },
    {
      "id": 298227256,
      "created_at": "2023-08-28T21:00:00.000Z",
      "title": "sansli numaram 933185",
      "db_time": "2023-09-01T21:55:22.790Z"
    },
    {
      "id": 298227255,
      "created_at": "2023-08-28T21:00:00.000Z",
      "title": "sansli numaram 878416",
      "db_time": "2023-09-01T21:55:22.790Z"
    }
  ],
  "connectionTime": "1ms",
  "queryTime": "54ms"
}
*/

// struct for the response
type Response struct {
	Success        bool     `json:"success"`
	Result         []Result `json:"result"`
	ConnectionTime string   `json:"connectionTime"`
	QueryTime      string   `json:"queryTime"`
}

// struct for the result
type Result struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Db_time    time.Time `json:"db_time"`
}

// Gets the rows with the limit given in the parameter
func getRowsWithLimit(c *gin.Context) {
	startTime := time.Now()
	limit := c.Query("q")
	if limit == "" {
		c.JSON(400, gin.H{
			"success": "false",
			"message": "No query provided",
		})
		return
	}

	query := "SELECT id, created_at, title, CURRENT_TIMESTAMP  db_time FROM test_table ORDER BY id DESC LIMIT " + limit

	results, _, a := execQuery(query)

	result := Response{
		Success:        true,
		Result:         results,
		ConnectionTime: a,
		QueryTime:      time.Since(startTime).String(),
	}

	c.JSON(200, result)
}

// Gets the max id of the table
func getMaxId(c *gin.Context) {
	startTime := time.Now()
	tableName := c.Query("q")
	if tableName == "" {
		c.JSON(400, gin.H{
			"success": "false",
			"message": "No query provided",
		})
		return
	}

	query := "SELECT max(id) FROM " + tableName

	rows, execTime, a := execQuery(query)

	c.JSON(200, gin.H{
		"success":          "true",
		"endPointCallTime": time.Since(startTime).String(),
		"queryExecTime":    execTime,
		"databaseConnTime": a,
		"message":          rows,
	})
}

// Gets random row by id
func getRandomRowById(c *gin.Context) {
	startTime := time.Now()

	query := "SELECT id FROM test_table WHERE id = " + strconv.Itoa(rand.Intn(5))

	rows, execTime, a := execQuery(query)

	c.JSON(200, gin.H{
		"success":          "true",
		"endPointCallTime": time.Since(startTime).String(),
		"queryExecTime":    execTime,
		"databaseConnTime": a,
		"message":          rows,
	})
}

/*


func query(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(400, gin.H{
			"success": "false",
			"message": "No query provided",
		})
		return
	}
	fmt.Println("Query:", query)

	pool, _, err := dbConnect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	result, queryTime := poolQuery(query, pool)

	if err != nil {
		fmt.Println("Error inserting into test table:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
		"time":    queryTime,
		"message": result,
	})
}

func insertRandomId(c *gin.Context) {
	randomId := strconv.Itoa(rand.Intn(100))
	fmt.Println("Inserting ", randomId, " into test table... ")

	query := (fmt.Sprintf(`
		insert into test (id) values (%s) returning id;
	`, randomId))

	pool, _, err := dbConnect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	result, queryTime := poolQuery(query, pool)

	if err != nil {
		fmt.Println("Error inserting into test table:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
		"time":    queryTime,
		"message": result,
	})

}

func getCount(c *gin.Context) {
	query := "Select  id maxID, created_at, title, CURRENT_TIMESTAMP  db_time FROM test_table where id = (select max(id) from test_table)"
	_, conn, err := dbConnect()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	result, len, queryTime := count(query, conn)
	resultString := fmt.Sprintf("%v", result)

	if err != nil {
		fmt.Println("Error inserting into test table:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
		"time":    queryTime,
		"length":  len,
		"message": resultString,
	})
}
*/
