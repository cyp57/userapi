package apihelper

import (
	"net/http"

	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
	"github.com/gin-gonic/gin"
)

type IResponseHandler interface{
	SuccessResponse(*gin.Context,int,interface{},string)
	ErrResponse(*gin.Context,int,string)
	DataResponse(*gin.Context,int,interface{})
}


type (
	ResponseHandler struct {
		Data    interface{}    `json:"data"`
		Message string `json:"message"`
		Status bool `json:"status"`
		Error *ErrorRes `json:"error"`
	}

	ErrorRes struct {
		ErrCode int `json:"errCode"`
		ErrMsg string `json:"errMsg"`
	}
)

// var resp = &ResponseHandler{}
var ErrCode = 0

func SetErrorCode(errCode int) {
	ErrCode = errCode
}

func (r *ResponseHandler) SuccessResponse(c *gin.Context , statusCode int , data interface{}, message string)  {
	c.JSON(statusCode, &ResponseHandler{Data: data , Message: message,Status: true})
}

func (r *ResponseHandler) ErrResponse(c *gin.Context , statusCode int , message string)  {
	if ErrCode != 0 && message == "" {
	errMsg	:= cnst.GetErrMsg(ErrCode)
	message = errMsg
	}
	c.JSON(statusCode, &ResponseHandler{Status: false,Error: &ErrorRes{ErrMsg: message,ErrCode: ErrCode}})
}

func (r *ResponseHandler) DataResponse(c *gin.Context , statusCode int , data interface{}) {
	c.JSON(statusCode, data)
}

















func (r *ResponseHandler) StatusDataResponse(result model.Response, c *gin.Context) {
	if result.StatusCode == http.StatusOK {
		// 200
		c.JSON(http.StatusOK, result.Result)
	} else if result.StatusCode == http.StatusBadRequest {
		// 400
		c.JSON(http.StatusBadRequest, result.Result)
	} else {
		// 500
		c.JSON(http.StatusInternalServerError, result.Result)
	}
}

func (r *ResponseHandler) StatusSaveResponse(result model.Response, c *gin.Context) {
	if result.StatusCode == http.StatusOK {
		// 200
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": result.Message, "results": result.Result, "saveStatus": result.Status})
	} else if result.StatusCode == http.StatusBadRequest {
		// 400
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": result.Message, "results": result.Result, "saveStatus": result.Status})
	} else if result.StatusCode == http.StatusUnauthorized {
		//401
		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": http.StatusUnauthorized, "message": result.Message, "results": result.Result, "saveStatus": result.Status})
	} else {
		//500
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": http.StatusInternalServerError, "message": result.Message, "results": result.Result, "saveStatus": result.Status})
	}
}



