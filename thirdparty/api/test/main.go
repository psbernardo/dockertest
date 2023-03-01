package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test")
	server := gin.Default()

	//GCP Kubernetes deployments health readiness probe checks
	server.GET("health", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	// Elastic Beanstalk forwards requests to port 5000
	server.Run(":8000")
}
