package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// Middleware 1- Request Validation
	r.Use(validateRequest)

	// Middleware 2- Request Tracing
	r.Use(traceRequest)

	// Middleware 3- Security
	r.Use(secureRequest)

	// Middleware 4- Authentication
	r.Use(authenticateRequest)

	r.POST("/login", login)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server")
	}
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login Successful"})
}

func validateRequest(c *gin.Context) {
	fmt.Println("Validating Request")
	c.Next()
}

func traceRequest(c *gin.Context) {
	fmt.Println("Tracing Request")
	beforeRequest(c)
	c.Next()
	afterRequest(c)

}

func authenticateRequest(c *gin.Context) {
	fmt.Println("Authenticating Request")
	c.Next()
}

func secureRequest(c *gin.Context) {
	fmt.Println("Securing Request")
	c.Next()
}

func beforeRequest(c *gin.Context) {
	start := time.Now()

	fmt.Printf("Started %s %s\n", c.Request.Method, c.Request.URL.Path)
	c.Set("startTime", start)
}

func afterRequest(c *gin.Context) {
	startTime, exists := c.Get("startTime")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time))
	fmt.Printf("Completes %s %s in %v\n", c.Request.Method, c.Request.URL.Path, duration)
}

// Example from : https://medium.com/@ansujain/mastering-middleware-in-go-tips-tricks-and-real-world-use-cases-79215e72b4a8
