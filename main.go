package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		// handle websocket requests
	})

	port := os.Getenv("PORT")
	r.Run(port)
}
