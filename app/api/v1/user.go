package v1

import (
	
	"net/http"

	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/utils"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	CreateUser(*gin.Context)
	CreateAdmin(*gin.Context)
	EditUser(*gin.Context)
	PatchUser(*gin.Context)
	GetUser(*gin.Context)
	GetUserList(*gin.Context)
	DeleteUser(*gin.Context)
	ForgotPassword(*gin.Context)
	ChangePassword(*gin.Context)
}

type UserApi struct{}

func (u *UserApi) CreateAdmin(c *gin.Context) {
	var jsonbody model.RegistrationInfo

	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		fusionAppId := utils.GetYaml(cnst.FusionAppId)
		result, err := userCtrlV1.CreateUser(&jsonbody, fusionAppId,true)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.SignupOk)
			return
		}
	}

}

func (u *UserApi) CreateUser(c *gin.Context) {
	var jsonbody model.RegistrationInfo

	if err := c.ShouldBindJSON(&jsonbody); err != nil {

		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		fusionAppId := utils.GetYaml(cnst.FusionAppId)
		result, err := userCtrlV1.CreateUser(&jsonbody, fusionAppId)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.SignupOk)
			return
		}
	}

}

func (u *UserApi) EditUser(c *gin.Context) {
	var jsonbody model.UserInfo
	uuid := c.Param("uuid")

	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		result, err := userCtrlV1.EditUser(uuid, &jsonbody)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.UpdateSuccess)
			return
		}
	}

}

func (u *UserApi) PatchUser(c *gin.Context) {
	jsonbody := make(map[string]interface{})
	uuid := c.Param("uuid")
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		result, err := userCtrlV1.PatchUser(uuid, jsonbody)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.UpdateSuccess)
			return
		}
	}
}

// list
func (u *UserApi) GetUserList(c *gin.Context) {

	filterInit := []string{"search", "limit", "page", "sort", "sortkey", "uuid"}
	filter := utils.CreateReqFilter(c, filterInit)

	result, err := userCtrlV1.GetUserList(filter)
	if err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
	} else {
		resp.DataResponse(c, http.StatusOK, &result)
	}
}

func (u *UserApi) GetUser(c *gin.Context) {
	uuid := c.Param("uuid")

	if !utils.IsEmptyString(uuid) {
		result, err := userCtrlV1.GetUserInfo(uuid)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		} else {
			resp.DataResponse(c, http.StatusOK, &result)
		}
	} else {
		resp.ErrResponse(c, http.StatusBadRequest, cnst.ErrReqPathParamUuid)
	}
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")

	if !utils.IsEmptyString(uuid) {
		result, err := userCtrlV1.DeleteUser(uuid)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		} else {
			resp.SuccessResponse(c, http.StatusOK, &result, cnst.DeleteSuccess)
		}
	} else {
		resp.ErrResponse(c, http.StatusBadRequest, cnst.ErrReqPathParamUuid)
	}
}

func (u *UserApi) ForgotPassword(c *gin.Context) {
	var jsonbody model.ForgotPasswordInfo

	if err := c.ShouldBindJSON(&jsonbody); err != nil {

		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		fusionAppId := utils.GetYaml(cnst.FusionAppId)
		result, err := userCtrlV1.ForgotPassword(&jsonbody, fusionAppId)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.ForgotPasswordSuccess)
			return
		}
	}

}
func (u *UserApi) ChangePassword(c *gin.Context) {
	var jsonbody model.ChangePasswordInfo

	tokenUuid := c.GetString("userId")
	uuid := c.Param("uuid")
	if uuid != tokenUuid {
		resp.ErrResponse(c, http.StatusBadRequest, cnst.ErrValidUuid)
		return
	}

	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {

		fusionAppId := utils.GetYaml(cnst.FusionAppId)
		result, err := userCtrlV1.ChangePassword(uuid, &jsonbody, fusionAppId)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.ChangePasswordSuccess)
			return
		}
	}

}
