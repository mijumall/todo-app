package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const listenAddr = "0.0.0.0:8002"

var URL = "postgres://root@crdb-1:29000/app"

func runRouter() {
	app := gin.Default()
	router := app.Group("/v1")

	isDev := os.Getenv("MODE") == "dev"
	if isDev {
		gin.SetMode(gin.DebugMode)
		router.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST"},
		}))
		URL = "postgres://root@cockroach:29000/app"
	} else {
		gin.SetMode(gin.ReleaseMode)
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"https://todo.anchorboard.net"},
			AllowMethods: []string{"GET", "POST"},
		}))
	}

	router.GET("/", func(c *gin.Context) { rootHandler(c) })
	router.GET("/read", readHandler)
	router.POST("/add", addHandler)
	router.POST("/delete", deleteHandler)

	app.Run(listenAddr)
}
