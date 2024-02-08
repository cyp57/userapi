package middleware

import (
	"bytes"
	"fmt"
	"net/http"

	"time"

	response "github.com/cyp57/user-api/app/api-helper"
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/pkg/logger"

	"github.com/gin-gonic/gin"
)


var resp = &response.ResponseHandler{}

type MiddlewareHandler struct{}

// Custom response writer to intercept response data
func (m *MiddlewareHandler) InterceptLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		log := &logger.LoggerObj{
			Time:   time.Now().Local().Format("2006-01-02 15:04:05"),
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.RequestURI,
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

func (m *MiddlewareHandler) ValidateToken(c *gin.Context) {
	
	if len(c.Request.Header.Get("token")) != 0 {
		token := c.Request.Header.Get("token")
		fmt.Println("token = = ",token)
		decode, err := new(fusionauth.Fusionauth).ValidateToken(token)
		if err != nil || decode.StatusCode == http.StatusUnauthorized{
			resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrTokenExpired)
			c.Abort()
		}
		fmt.Println("decode = =", decode)
		uuid := decode.Jwt.Sub
		// c.Set()
		c.Set("userId",uuid)
		c.Next()
	} else {
		resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrReqToken)
		c.Abort()
	}
}

// / cors gin
// corsMiddleware handles CORS headers
func (m *MiddlewareHandler) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// // validate role
// func AuthorizeRole(expectRole ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		user, exists := c.Get("user")
// 		if !exists {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}

// 		userRoles, ok := user.(*models.User).Roles
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}

// 		// Check if the user has the required role
// 		if !hasRole(userRoles, role) {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
// 			return
// 		}

// 		c.Next()
// 	}
// }

func hasRole(userRoles []string, role string) bool {
	for _, r := range userRoles {
		if r == role {
			return true
		}
	}
	return false
}
