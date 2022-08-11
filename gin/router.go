package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func runRouter(conn *pgx.Conn) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) { rootHandler(c) })
	router.GET("/read", func(c *gin.Context) { readHandler(c, conn) })
	router.POST("/add", func(c *gin.Context) { addHandler(c, conn) })
	router.DELETE("/delete", func(c *gin.Context) { deleteHandler(c, conn) })

	router.Run("0.0.0.0:8002")
}
