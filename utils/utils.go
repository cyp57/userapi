package utils

import (
	"encoding/json"
	"fmt"
	"time"

	oid "github.com/coolbed/mgo-oid"
	"github.com/gin-gonic/gin"
)

func Debug(data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("json.MarshalIndent err:", err.Error())
	}
	fmt.Println(string(bytes))
}

func Output(data interface{}) []byte {
	bytes, _ := json.Marshal(data)
	return bytes
}

func containInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func GenerateOid(prefix string) string {
	objId := oid.NewOID()

	return prefix + objId.String()
}

func IsEmptyString(s string) bool {
	return s == ""
}

func GetCurrentTime() (time.Time, error) {
	// Get the current time
	currentTime := time.Now()

	// thaizone, err := time.LoadLocation("Asia/Bangkok")
	// if err != nil {
	// 	return time.Time{}, err
	// }
	timeString := currentTime.Local().Format(time.RFC3339)
	// timeString := currentTime.In(thaizone).Format(time.RFC3339)

	parsedTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}, err
	}

	// Return the parsed time
	return parsedTime, nil
}

func CreateReqFilter(c *gin.Context, arrStr []string ) map[string]interface{} {
	filterMap := make(map[string]interface{})
	for _, key := range arrStr {
		if c.Request.URL.Query().Get(key) != "" {
			filterMap[key] = c.Request.URL.Query().Get(key)
		}
	}

	return filterMap
}
