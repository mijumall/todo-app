package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"root": "Hello under world",
	})
}

func readHandler(c *gin.Context) {
	var data []Data

	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer conn.Close(context.Background())

	sql := `SELECT id, content, created_at FROM todo`
	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	for rows.Next() {
		var d Data
		rows.Scan(&d.Id, &d.Content, &d.CreatedAt)
		data = append(data, d)
	}

	c.IndentedJSON(http.StatusOK, data)
}

func addHandler(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer conn.Close(context.Background())

	// Body will be: {"content": "Some todo text"}
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	sql := `INSERT INTO todo (content, created_at) VALUES ($1, $2)`
	tag, err := conn.Exec(context.Background(), sql, body["content"], time.Now().Format(time.RFC3339Nano))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": false,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"result":        true,
		"rows_affected": tag.RowsAffected(),
	})
}

func deleteHandler(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer conn.Close(context.Background())

	id := c.Query("id")

	sql := `DELETE FROM todo WHERE id = $1`
	_, err = conn.Exec(context.Background(), sql, id)
	if err != nil {
		c.JSON(400, gin.H{
			"result": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": true,
	})
}
