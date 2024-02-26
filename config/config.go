package config

import (
	"log"

	"github.com/cyp57/userapi/cnst"
	lrlog "github.com/cyp57/userapi/pkg/logrus"
	"github.com/cyp57/userapi/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig(envPath, yamlPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalln(err)
	}

	lrlog.InitLogrus()
	mode := utils.GetEnv(cnst.Mode)

	var viperYaml = viper.New()
	viperYaml.SetConfigName(mode)

	viperYaml.SetConfigType("yaml")
	viperYaml.AddConfigPath(yamlPath)

	err = viperYaml.ReadInConfig()
	if err != nil {
		// 	log.Fatalln(err) or use build in log
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	} else {
		utils.SetViperYaml(viperYaml)
	}
}
