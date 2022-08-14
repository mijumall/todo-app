package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func runRouter(conn *pgx.Conn) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8001",
			"https://localhost",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.GET("/", func(c *gin.Context) { rootHandler(c) })
	router.GET("/read", func(c *gin.Context) { readHandler(c, conn) })
	router.POST("/add", func(c *gin.Context) { addHandler(c, conn) })
	router.DELETE("/delete", func(c *gin.Context) { deleteHandler(c, conn) })

	router.Run("0.0.0.0:8002")
}
