package v1

import (
	"fmt"
	"net/http"

	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/utils"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	CreateUser(*gin.Context)
	EditUser(*gin.Context)
	GetUser(*gin.Context)
}

type UserApi struct{}

func (u *UserApi) CreateUser(c *gin.Context) {
	var jsonbody model.RegistrationInfo

	///1.check
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		fmt.Println("ShouldBindJSON")
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		fusionAppId := utils.GetYaml("FusionAppId")
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
	///1.check
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		fmt.Println("ShouldBindJSON")
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

func (u *UserApi) PatchPerson(c *gin.Context) {
	jsonbody := make(map[string]interface{})

	///1.check
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		// log.Fatalln(err.Error())
		// responseHandler.StatusSaveResponse(structs.Jsonresponse{Message: err.Error(), StatusCode: http.StatusBadRequest}, c)
		// return
	} else {

	}

	c.JSON(200, jsonbody)

}

// list
func (u *UserApi) ListPerson(c *gin.Context) {

	c.JSON(200, nil)
}

// get obj
func (u *UserApi) GetUser(c *gin.Context) {
	uuid := c.Param("uuid")

	if !utils.IsEmptyString(uuid) {
		result, err := userCtrlV1.GetUserInfo(uuid)
		if err != nil {
			resp.ErrResponse(c, http.StatusInternalServerError, err.Error())
		} else {
			resp.DataResponse(c, http.StatusOK, &result)
		}
	} else {
		resp.ErrResponse(c, http.StatusBadRequest, cnst.ErrReqPathParamUuid)
	}
}
