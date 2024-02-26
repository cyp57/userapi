package middleware

import (
	"bytes"

	"net/http"

	"time"

	response "github.com/cyp57/userapi/app/response"
	"github.com/cyp57/userapi/cnst"
	"github.com/cyp57/userapi/pkg/fusionauth"
	"github.com/cyp57/userapi/pkg/logger"
	"github.com/cyp57/userapi/utils"

	"github.com/gin-gonic/gin"
)

var resp = response.Response()

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
		// log.Print()  print log only.
		log.Print().Save() //  print and save log to \assets\log
	}
}

func (m *MiddlewareHandler) ValidateToken(c *gin.Context) {

	if len(c.Request.Header.Get("token")) != 0 {
		token := c.Request.Header.Get("token")

		var fusionObj fusionauth.Fusionauth
		var appId = utils.GetYaml(cnst.FusionAppId)
		fusionObj.SetApplicationId(appId)

		decode, err := fusionObj.ValidateToken(token)
		if err != nil || decode.StatusCode == http.StatusUnauthorized {
			resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrTokenExpired)
			// c.Abort()
			return
		}

		uuid := decode.Jwt.Sub
		userInfo, err := fusionObj.GetUserRegistration(uuid)
		utils.Debug("MiddlewareHandler.ValidateToken :")
		utils.Debug(userInfo)

		if err != nil {
			resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrReqRole)
			// c.Abort()
			return
		}

		c.Set("userId", uuid)
		c.Set("roles", userInfo.Registration.Roles)
		c.Next()
	} else {
		resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrReqToken)
		c.Abort()
	}
}

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

func (m *MiddlewareHandler) AuthorizeRole(expectRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(expectRole) > 0 {
			userRoles, exists := c.Get("roles")

			if !exists {
				resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrAuthorizeRole)
				return
			}

			roles, ok := userRoles.([]string)
			if !ok {
				resp.ErrResponse(c, http.StatusUnauthorized, cnst.ErrInvalidRole)
				return
			}

			var result bool
			for _, check := range expectRole {

				result = hasRole(roles, check)
				if result {
					break
				}
			}

			if !result {
				resp.ErrResponse(c, http.StatusForbidden, cnst.ErrForbidden)
				return
			}
			c.Next()
		}
	}
}

func hasRole(userRoles []string, role string) bool {
	for _, r := range userRoles {
		if r == role {
			return true
		}
	}
	return false
}
