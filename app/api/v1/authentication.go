package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// response "github.com/cyp57/userapi/app/api-helper"
	"github.com/cyp57/userapi/cnst"
	"github.com/cyp57/userapi/model"
	"github.com/cyp57/userapi/utils"
)

type IAuthentication interface {
	Login(c *gin.Context)
	Refresh(c *gin.Context)
	LogOut(c *gin.Context)
}

type AuthenticationApi struct{}

func (a *AuthenticationApi) Login(c *gin.Context) {
	var loginObj model.LoginInfo
	if err := c.ShouldBindJSON(&loginObj); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		var appId = utils.GetYaml(cnst.FusionAppId)
		result, err := authCtrlV1.Login(&loginObj, appId)

		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.LoginSuccess)
			return
		}

	}
}

func (a *AuthenticationApi) Refresh(c *gin.Context) {
	var jsonbody model.RefreshJwt
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		var appId = utils.GetYaml(cnst.FusionAppId)
		result, err := authCtrlV1.RefreshJwt(&jsonbody, appId)

		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.RequestSuccess)
			return
		}

	}
}

func (a *AuthenticationApi) LogOut(c *gin.Context) {
	var jsonbody model.LogOutInfo
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		result, err := authCtrlV1.LogOut(&jsonbody)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.RequestSuccess)
			return
		}

	}
}
