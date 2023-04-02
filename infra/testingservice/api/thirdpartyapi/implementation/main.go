package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}

func main() {
	fmt.Println("test")
	server := gin.Default()

	//GCP Kubernetes deployments health readiness probe checks
	server.GET("health", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	//GCP Kubernetes deployments health readiness probe checks
	server.GET("person/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, Person{
			ID:       3,
			Name:     "Patrick",
			LastName: "Bernardo",
			Age:      28,
		})
	})

	// Elastic Beanstalk forwards requests to port 5000
	server.Run(":8000")
}
