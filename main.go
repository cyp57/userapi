package main

import (

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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	FileEnv = "config/.env"
	FileIni = "config/app.ini"
)

func main() {
	godotenv.Load(FileEnv)
	cnst.InitErr()
	initService()
}

func initService() {
	lslog.InitLogrus() //checked
	initConfigMode()   //checked

	initSetting() //checked

	mongodb.MongoDbConnect() //checked
	initFusionAuth()         //checked
	route.InitRoute()  //checked
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
	viperYaml.SetConfigName(env)

	viperYaml.SetConfigType("yaml")
	viperYaml.AddConfigPath("config")

	// เริ่มการค้นหาไฟล์ Config และอ่านไฟล์
	err := viperYaml.ReadInConfig()
	if err != nil {
		// log.Fatalln(cnst.ErrParseConfigYaml)
		ls := &lslog.LslogObj{Data: nil, Txt: cnst.ErrParseConfigYaml, Level: logrus.FatalLevel}
		ls.Print()
	} else {
		utils.SetViperYaml(viperYaml)
	}
}

func initSetting() {
	init := new(setting.ApiSetting)
	init.Setup(FileIni)
}

func initFusionAuth() {
	var host = utils.GetYaml(cnst.FusionHost)
	var apiKey = utils.GetYaml(cnst.FusionAPIKey)
	var httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	var baseURL, err = url.Parse(host)
	if err != nil {
		ls := &lslog.LslogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}
	// Construct a new FusionAuth Client
	Auth := fusionauth.NewClient(httpClient, baseURL, apiKey)

	// for production code, don't ignore the error!
	_, err = Auth.RetrieveTenants()
	if err != nil {
		ls := &lslog.LslogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}

	new(fusionauthPkg.Fusionauth).InitConnection(Auth)
}
