package main

import (
	"log"
	"lp-assistant/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register analytics endpoint
	r.GET("/analytics", handlers.AnalyticsHandler)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
