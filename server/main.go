package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("no listening port provided")
	}

	fmt.Printf("Listening at :%v/ping ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
