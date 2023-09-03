package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success        bool     `json:"success"`
	Result         []Result `json:"result"`
	ConnectionTime string   `json:"connectionTime"`
	QueryTime      string   `json:"queryTime"`
}
type Response2 struct {
	Success        bool    `json:"success"`
	Result         Result2 `json:"result"`
	ConnectionTime string  `json:"connectionTime"`
	QueryTime      string  `json:"queryTime"`
}

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
	tableName := c.Query("q")
	if tableName == "" {
		c.JSON(400, gin.H{
			"success": "false",
			"message": "No query provided",
		})
		return
	}

	rows, execTime, a := maxId(tableName)

	result := Response2{
		Success:        true,
		Result:         rows,
		ConnectionTime: a,
		QueryTime:      execTime,
	}

	c.JSON(200, result)

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
