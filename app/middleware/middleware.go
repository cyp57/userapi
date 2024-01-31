package middleware

import (
	"bytes"

	"time"

	"github.com/cyp57/user-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

type MiddlewareHandler struct{}

// Custom response writer to intercept response data
func (m *MiddlewareHandler) InterceptLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		log := &logger.LoggerObj{
			Time:   time.Now().Local().Format("2006-01-02 15:04:05"),
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.RequestURI,
			// StatusCode: c.Request.Response.StatusCode,
		}
		
		log.SetQuery(c)
		log.SetBody(c)

		// Create a buffer to capture the response
		buff := &bytes.Buffer{}
		rw := &logger.ResponseWriter{Body: buff, ResponseWriter: c.Writer}

		// Set the intercepted writer to the context
		c.Writer = rw

		// Process the request
		c.Next()

		log.SetResponse(c, rw)
		// log.Print()
		log.Print().Save()
	}
}

// Custom response writer to intercept response data
// type responseWriter struct {
// 	gin.ResponseWriter
// 	body   *bytes.Buffer
// 	status int
// }

// In this example, the responseWriter now also stores the status code. The middleware checks if the response is JSON based on the "Content-Type" header and then logs the JSON response, body, and status code accordingly.

// Make sure to adjust the middleware and route logic according to your specific use case and requirements.



/// cors gin
// corsMiddleware handles CORS headers

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent)
// 			return
// 		}

// 		c.Next()
// 	}
// }