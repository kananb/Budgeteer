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

	r.Static("/", "/app/client/build")
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("no listening port provided")
	}

	fmt.Printf("Listening at :%v ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
