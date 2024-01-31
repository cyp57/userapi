package setting

import (
	"github.com/cyp57/user-api/cnst"
	lslog "github.com/cyp57/user-api/pkg/logrus"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type ApiSetting struct{}

type ApiGroup struct {
	AuthenGroup       string
	LoggingGroup      string
	LogDataGroup      string
	ImageServiceGroup string
}

var ApiGroupSetting = &ApiGroup{}

type ApiEndpoint struct {
	AuthenService       string
	LoggingService      string
	LogDataService      string
	ImageUrl            string
	ImageUpload         string
	ImageDeleteEndPoint string
}

var ApiEndpointSetting = &ApiEndpoint{}

type Collection struct {
	LogService      string
	TemplateService string
	User string
}

var CollectionSetting = &Collection{}

var cfg *ini.File

// Setup initialize the configuration instance
func (a *ApiSetting) Setup(iniPath string) {

	var err error
	cfg, err = ini.Load(iniPath)
	if err != nil {
		ls := &lslog.LslogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
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
		ls := &lslog.LslogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}
}
