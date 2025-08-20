package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword:= os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")



	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"dbUser": dbUser,
			"dbPassword": dbPassword,
			"dbName": dbName,
		})
	})
	r.GET("/batyr", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"dbUser": dbUser,
			"dbPassword": dbPassword,
			"dbName": dbName,
		})
	})
	r.Run()
}
