package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	lrlog "github.com/cyp57/user-api/pkg/logrus"
	"github.com/cyp57/user-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body   *bytes.Buffer
	status int
}

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(c *gin.Context)
	SetBody(c *gin.Context)
	SetResponse(c *gin.Context, rw *ResponseWriter)
}

type LoggerObj struct {
	Time       string      `json:"time"`
	Ip         string      `json:"ip"`
	Path       string      `json:"path"`
	Method     string      `json:"method"`
	StatusCode int         `json:"statusCode"`
	Query      interface{} `json:"query"`
	Body       interface{} `json:"body"`
	Response   interface{} `json:"response"`
}

func (l *LoggerObj) Print() ILogger {
	utils.Debug(l)
	return l
}
func (l *LoggerObj) Save() {
	data := utils.Output(l)
	filename := fmt.Sprintf("./assets/log/logger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}
func (l *LoggerObj) SetQuery(c *gin.Context) {
	// Retrieve all query parameters as a map
	queryParams := c.Request.URL.Query()

	// Log all query parameters
	log.Printf("Received query parameters: %v", queryParams)

	l.Query = queryParams

}
func (l *LoggerObj) SetBody(c *gin.Context) {
	// var body interface{}
	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut {
		// Read the request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.ErrorLevel}
			ls.Print()
			return
		}
		// Log the request body
		log.Printf("Received request body: %s", body)
		// Declare an empty interface{}
		var result interface{}

		// Unmarshal the byte slice into the empty interface
		err = json.Unmarshal(body, &result)
		if err != nil {
			ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.ErrorLevel}
			ls.Print()
			return
		}

		l.Body = result
		// Rewind the request body so it can be read again if needed
		c.Request.Body = io.NopCloser(strings.NewReader(string(body)))
	}
}
func (l *LoggerObj) SetResponse(c *gin.Context, rw *ResponseWriter) {

	// Store the response writer to intercept the response

	var resMap interface{}
	err := json.Unmarshal(rw.Body.Bytes(), &resMap)
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.ErrorLevel}
		ls.Print()
		return
	}

	l.StatusCode = rw.Status()
	l.Response = resMap
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *ResponseWriter) WriteString(s string) (int, error) {
	rw.Body.WriteString(s)
	return rw.ResponseWriter.WriteString(s)
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Status() int {
	if rw.status == 0 {
		return http.StatusOK
	}
	return rw.status
}
