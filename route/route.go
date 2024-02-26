package route

import (
	"path"

	v1 "github.com/cyp57/userapi/app/api/v1"
	"github.com/cyp57/userapi/app/middleware"
	"github.com/cyp57/userapi/cnst"
	"github.com/cyp57/userapi/utils"
	"github.com/easonlin404/limit"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {

	router := gin.Default()

	httpRequestLimit := utils.GetYamlInt(cnst.HttpRequestLimit)
	httpport := utils.GetYaml(cnst.HTTPPort)

	if httpRequestLimit != 0 {
		router.Use(limit.Limit(httpRequestLimit))
	}

	router.Use(new(middleware.MiddlewareHandler).CorsMiddleware())
	setRoute(router)

	router.Run(":" + httpport)
	return router
}

func setRoute(router *gin.Engine) {

	api := v1.InitApiUserImpl()

	middleHandler := new(middleware.MiddlewareHandler)

	serviceName := utils.GetYaml(cnst.ServiceName)
	router.GET(serviceName, root)
	v1 := router.Group(serviceName)
	v1.Use(middleHandler.InterceptLog())
	{ // non require token group
		nonAuthGroup := v1.Group(path.Join("user"))
		nonAuthGroup.POST("/login", api.Login)
		nonAuthGroup.POST("/logout", api.LogOut)
		nonAuthGroup.POST("/refresh/token", api.Refresh)
		nonAuthGroup.POST("/signup", api.CreateUser)
		nonAuthGroup.POST("/forgot/password", api.ForgotPassword)
	}
	{ // require token group
		reqAuthGroup := v1.Group(path.Join("user"))
		reqAuthGroup.Use(middleHandler.ValidateToken)
		reqAuthGroup.POST("/admin/signup", middleHandler.AuthorizeRole("admin"), api.CreateAdmin)       /// admin
		reqAuthGroup.PUT("/:uuid", middleHandler.AuthorizeRole("admin", "customer"), api.EditUser)      //// customer ,admin
		reqAuthGroup.PATCH("/:uuid", middleHandler.AuthorizeRole("admin", "customer"), api.PatchUser)   //// customer ,admin
		reqAuthGroup.GET("/:uuid", middleHandler.AuthorizeRole("admin", "customer"), api.GetUser)       // customer admin
		reqAuthGroup.GET("/", middleHandler.AuthorizeRole("admin"), api.GetUserList)                    // list for admin
		reqAuthGroup.DELETE("/:uuid", middleHandler.AuthorizeRole("admin", "customer"), api.DeleteUser) // customer admin

		reqAuthGroup.PUT("/change/password/:uuid", middleHandler.AuthorizeRole("admin", "customer"), api.ChangePassword) //// customer ,admin

	}

}

// for health check
func root(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
