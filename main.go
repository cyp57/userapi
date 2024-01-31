package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	// "os"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/cyp57/user-api/cnst"

	fusionauthPkg "github.com/cyp57/user-api/pkg/fusionauth"
	lslog "github.com/cyp57/user-api/pkg/logrus"
	"github.com/cyp57/user-api/pkg/mongodb"
	"github.com/cyp57/user-api/route"
	"github.com/cyp57/user-api/setting"
	"github.com/cyp57/user-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	FileEnv = ".env"
	FileIni = "config/app.ini"
)

func main() {
	godotenv.Load(FileEnv)
	cnst.InitErr()
	initService()
}

func initService() {
	initConfigMode()
	lslog.InitLogrus()
	initSetting()
	// cb := &lslog.LslogObj{Data: nil, Txt: "print LA FAt", Level: logrus.FatalLevel}
	// cb.LogrusPrint()

	mongodb.MongoDbConnect()
	initFusionAuth()
	route.InitRoute()
}

func initConfigMode() {
	mode := utils.GetEnv(cnst.Mode)
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
		//In release mode, Gin disables debugging features and provides better performance compared to the default development mode.
	}
	InitConfigYaml(mode)
}

func InitConfigYaml(env string) {

	var viperYaml = viper.New()
	viperYaml.SetConfigName(env) // ชื่อไฟล์ Config

	viperYaml.SetConfigType("yaml")
	viperYaml.AddConfigPath("config")

	// เริ่มการค้นหาไฟล์ Config และอ่านไฟล์
	err := viperYaml.ReadInConfig()
	if err != nil {
		log.Fatalln("error on parsing configuration filet")
	} else {
		utils.SetViperYaml(viperYaml)
	}
}

//////////////////////////////
// func initService() {
// 	initLogging()
// 	initConfigMode()
// 	initIni()
// 	initOtherService()
// 	initRouter()
// }

// func initOtherService() {
// 	authenservice.GetToken()
// 	initFusionAuth()
// }

// func initLogging() {

// 	var logConfig structs.LogConfiguration
// 	logConfig.Server = utils.GetEnv(cnst.ServerLogging)
// 	logConfig.ServicePath = path.Join(setting.ApiGroupSetting.LoggingGroup, setting.ApiEndpointSetting.LoggingService)
// 	logConfig.AppId = utils.GetEnv(cnst.AppId)
// 	logConfig.AppName = utils.GetEnv(cnst.AppName)
// 	logConfig.Level = utils.GetEnv(cnst.LogLevel)
// 	logConfig.OnServerLog = utils.GetEnvBool(cnst.OnServerLog)
// 	logging.InitLog(logConfig)
// }

func initSetting() {
	init := new(setting.ApiSetting)
	init.Setup(FileIni)
}

// func initConfigMode() {
// 	mode := utils.GetEnv(cnst.Mode)
// 	if mode == "prod" {
// 		rlimit.Setup()
// 		gin.SetMode(gin.ReleaseMode)
// 	}
// 	conf.InitConfigYaml(mode)
// }

func initFusionAuth() {
	var host = utils.GetYaml(cnst.FusionHost)
	var apiKey = utils.GetYaml(cnst.FusionAPIKey)
	var httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	var baseURL, err = url.Parse(host)
	if err != nil {
		log.Fatalln(err)
	}
	// Construct a new FusionAuth Client
	Auth := fusionauth.NewClient(httpClient, baseURL, apiKey)
	fmt.Println("Auth= = ", Auth)
	fmt.Println("Auth*= = ", *Auth)
 // for production code, don't ignore the error!
 tenantResponse, err := Auth.RetrieveTenants() 
    if err != nil {
		log.Fatalln(err.Error())
	}
 fmt.Print(len(tenantResponse.Tenants))
	

	new(fusionauthPkg.Fusionauth).InitConnection(Auth)
}



// func initRouter() {
// 	routersInit := routers.InitRoute()      ///initial path from gin
// 	port := utils.GetYaml(appCnst.HTTPPort) ///get httpport from .env
// 	routersInit.Run(":" + port)
// }
