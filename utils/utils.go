package utils

import (
	"encoding/json"
	"fmt"
	"time"

	oid "github.com/coolbed/mgo-oid"
)

func Debug(data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "\t")
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

	thaizone, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return time.Time{}, err
	}

	timeString := currentTime.In(thaizone).Format(time.RFC3339)

	parsedTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}, err
	}

	// Return the parsed time
	return parsedTime, nil
}
