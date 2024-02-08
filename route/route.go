package route

import (
	"path"

	v1 "github.com/cyp57/user-api/app/api/v1"
	"github.com/cyp57/user-api/app/middleware"
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/utils"
	"github.com/easonlin404/limit"
	"github.com/gin-gonic/gin"
)

// InitRoute ..
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
	{
		nonAuthGroup := v1.Group(path.Join("user"))
		nonAuthGroup.POST("/login", api.Login)
		nonAuthGroup.POST("/signup", api.CreateUser)
	}
	{
		reqAuthGroup := v1.Group(path.Join("user"))
		reqAuthGroup.Use(middleHandler.ValidateToken)
		reqAuthGroup.PUT("/:uuid", api.EditUser) //// customer ,admin
		reqAuthGroup.PATCH("/:uuid", api.PatchUser) //// customer ,admin
	    reqAuthGroup.GET("/:uuid", api.GetUser) // customer admin
		reqAuthGroup.GET("/", api.GetUserList) // list for admin 
	}

	// nonAuthGroup := router.Group(serviceName)
	// nonAuthGroup.Use(middleHandler.InterceptLog())
	// {
	// 	nonAuthGroup.POST("/login", api.Login)
	// 	nonAuthGroup.POST("/signup", api.CreateUser)
	// }

	// reqAuthGroup := router.Group(serviceName)
	// reqAuthGroup.Use(middleHandler.InterceptLog())
	// reqAuthGroup.Use(middleHandler.ValidateToken)
	// {
	// 	// v1.POST("/admin/signup", api.CreateUser) for admin only

	// 	reqAuthGroup.PUT("/user/:id", api.EditUser) //// customer ,admin
	// 	// v1.PATCH("/user", api.EditUser)   //// customer ,admin
	// 	// v1.get("/user", api.EditUser)  list  / admin
	// 	// v1.get("/user", api.EditUser)  by id // customer admin
	// 	// v1.PUT("/user", api.EditUser)  delete  //admin

	// 	// v1.PUT forgot pass  customer admin
	// 	// v1.PUT change pass  customer admin

	// 	// path := new(pathRoute.Path)
	// 	// pathGroup := v1.Group("/" + setting.SetupSetting.RouterGroup)
	// 	// path.PathRoute(pathGroup)
	// }

}

func root(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
