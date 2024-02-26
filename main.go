package main

import (
	"github.com/cyp57/userapi/cnst"
	"github.com/cyp57/userapi/config"
	fusionauthPkg "github.com/cyp57/userapi/pkg/fusionauth"
	lrlog "github.com/cyp57/userapi/pkg/logrus"
	"github.com/cyp57/userapi/pkg/mongodb"
	"github.com/cyp57/userapi/route"
	"github.com/cyp57/userapi/setting"
	"github.com/cyp57/userapi/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	PathEnv  = "config/.env"
	PathIni  = "config/app.ini"
	PathYaml = "config"
)

func main() {
	cnst.InitErr()
	initService()
}

func initService() {

	config.LoadConfig(PathEnv, PathYaml)
	setting.InitIni(PathIni)
	mongodb.MongoDbConnect()
	fusionauthPkg.InitFusionAuth()
	route.InitRoute()

}

func InitConfigMode() {
	mode := utils.GetEnv(cnst.Mode)
	InitConfigYaml(mode)
}

func InitConfigYaml(env string) {

	var viperYaml = viper.New()
	viperYaml.SetConfigName(env)

	viperYaml.SetConfigType("yaml")
	viperYaml.AddConfigPath("config")

	err := viperYaml.ReadInConfig()
	if err != nil {
		// log.Fatalln(cnst.ErrParseConfigYaml)
		ls := &lrlog.LrlogObj{Data: nil, Txt: cnst.ErrParseConfigYaml, Level: logrus.FatalLevel}
		ls.Print()
	} else {
		utils.SetViperYaml(viperYaml)
	}
}
