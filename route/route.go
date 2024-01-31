package route

import (
	"net/http"

	v1 "github.com/cyp57/user-api/app/api/v1"
	"github.com/cyp57/user-api/app/middleware"
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/utils"
	"github.com/easonlin404/limit"
	"github.com/gin-gonic/gin"
)

// InitRoute ..
func InitRoute() *gin.Engine {

	// Creates a router without any middleware by default
	router := gin.New()
	httpRequestLimit := utils.GetYamlInt(cnst.HttpRequestLimit)
	httpport := utils.GetYaml(cnst.HTTPPort)

	if httpRequestLimit != 0 {
		router.Use(limit.Limit(httpRequestLimit))
	}

	//mode := utils.GetEnv(cnst.Mode)
	router.Use(gin.Logger())
	// router.Use(middlewares.GinBodyLogMiddleware(logging.Debug, mode))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// If use gin.Logger middlewares, it send duplicated request.
	setRoute(router)

	router.Run(":" + httpport) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return router
}

func setRoute(router *gin.Engine) {
	ua := new(v1.UserApi)
	auth := new(v1.AuthenticationApi)
	/// middleware check token
	middleHandler := new(middleware.MiddlewareHandler)

	//// get servicename
	serviceName := utils.GetYaml(cnst.ServiceName)
	router.GET(serviceName, root)
	v1 := router.Group(serviceName)
	v1.Use(middleHandler.InterceptLog())
	{
		
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/signup", ua.CreatePerson)
		v1.POST("/login", auth.Login)

		// path := new(pathRoute.Path)
		// pathGroup := v1.Group("/" + setting.SetupSetting.RouterGroup)
		// path.PathRoute(pathGroup)
	}

}

func root(c *gin.Context) { //{"message":"OK"}
	c.JSON(200, gin.H{"message": "OK"})
}
