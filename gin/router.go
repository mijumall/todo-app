package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func runRouter() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.GET("/", func(c *gin.Context) { rootHandler(c) })
	router.GET("/read", readHandler)
	router.POST("/add", addHandler)
	router.DELETE("/delete", deleteHandler)

	router.Run("0.0.0.0:8002")
}
