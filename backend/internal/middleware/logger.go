package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()

		fmt.Println("Incoming Request:")
		fmt.Println(c.Request.Method, c.Request.URL.Path)

		c.Next()

		duration := time.Since(startTime)

		fmt.Println("Request Completed In:", duration)
	}
}