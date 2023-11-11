package main

import (
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	FileEnv = ".env"
	FileIni = "conf/app.ini"
)

func main() {
	godotenv.Load(FileEnv)
	initService()
}

func initService() {
	initConfigMode()
}

func initConfigMode() {
	mode := utils.GetEnv(cnst.Mode)
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	InitConfigYaml(mode)
}

func InitConfigYaml(env string) {

	var viperYaml = viper.New()
	viperYaml.SetConfigName(env) // ชื่อไฟล์ Config

	viperYaml.SetConfigType("yaml")
	viperYaml.AddConfigPath("conf")

	// เริ่มการค้นหาไฟล์ Config และอ่านไฟล์
	err := viperYaml.ReadInConfig()
	if err != nil {
		//log.Fatal("error on parsing configuration file")
	} else {
		utils.SetViperYaml(viperYaml)
	}

}
