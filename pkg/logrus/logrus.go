package logrus

import (
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/utils"
	"github.com/sirupsen/logrus"
)

type LslogObj struct {
	Data  interface{}  `json:"data"`
	Txt   string       `json:"txt"`
	Level logrus.Level `json:"level"`
}

// init log logrus
func InitLogrus() {
	logrus.SetFormatter(&nested.Formatter{
		HideKeys: true,
		// FieldsOrder:     []string{"application", "module", "method"},
		TimestampFormat: time.RFC3339,
		ShowFullLevel:   true,
		NoFieldsSpace:   false,
		NoFieldsColors:  false,
		NoColors:        false,
		TrimMessages:    false,
	})

	logrus.Info("Log with RFC3339 timestamp format.")

	// Only log the warning severity or above.
	// logrus.SetLevel(logrus.DebugLevel)
	loglevel := utils.GetEnv(cnst.LogLevel)
	logrus.Info("Logrus level: ", loglevel)
	logrus.SetLevel(GetLogrusLevel(loglevel))
}

func (ls *LslogObj) Print() {
	switch ls.Level {
	case logrus.DebugLevel:
		logrus.WithFields(logrus.Fields{
			"json": ls.Data,
		}).Debug(ls.Txt)
	case logrus.InfoLevel:
		logrus.WithFields(logrus.Fields{
			"json": ls.Data,
		}).Info(ls.Txt)
	case logrus.ErrorLevel:
		logrus.WithFields(logrus.Fields{
			"json": ls.Data,
		}).Error(ls.Txt)
	case logrus.FatalLevel:
		logrus.WithFields(logrus.Fields{
			"json": ls.Data,
		}).Fatal(ls.Txt)
	default: // default as debug level
		logrus.WithFields(logrus.Fields{
			"json": ls.Data,
		}).Debug(ls.Txt)
	}

}

func GetLogrusLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "Debug":
		return logrus.DebugLevel
	case "Info":
		return logrus.InfoLevel
	case "Warning":
		return logrus.WarnLevel
	case "Error":
		return logrus.ErrorLevel
	case "Fatal":
		return logrus.FatalLevel
	}
	return logrus.DebugLevel

}

