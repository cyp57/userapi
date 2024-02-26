package setting

import (
	"github.com/cyp57/userapi/cnst"
	lrlog "github.com/cyp57/userapi/pkg/logrus"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type ApiGroup struct {
	UploadApiV1 string
}

var ApiGroupSetting = &ApiGroup{}

type ApiEndpoint struct {
	LogEndpoint    string
	UploadEndpoint string
}

var ApiEndpointSetting = &ApiEndpoint{}

type Collection struct {
	User string
}

var CollectionSetting = &Collection{}

var cfg *ini.File

// Setup initialize the configuration instance
func InitIni(iniPath string) {

	var err error
	cfg, err = ini.Load(iniPath)
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}

	mapTo(cnst.ApiGroup, ApiGroupSetting)
	mapTo(cnst.ApiEndpoint, ApiEndpointSetting)
	mapTo(cnst.Collection, CollectionSetting)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}
}
