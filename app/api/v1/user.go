package v1

import (
	"fmt"
	"net/http"

	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) CreatePerson(c *gin.Context) {
	var jsonbody model.UserInfo

	///1.check
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		// log.Fatalln(err.Error())
		// responseHandler.StatusSaveResponse(structs.Jsonresponse{Message: err.Error(), StatusCode: http.StatusBadRequest}, c)
		// return
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

func (u *UserApi) EditPerson(c *gin.Context) {
	var jsonbody model.UserInfo

	///1.check
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		// log.Fatalln(err.Error())
		// responseHandler.StatusSaveResponse(structs.Jsonresponse{Message: err.Error(), StatusCode: http.StatusBadRequest}, c)
		// return
	}

	c.JSON(200, jsonbody)

}

func (u *UserApi) PatchPerson(c *gin.Context) {
	jsonbody := make(map[string]interface{})

	///1.check
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		// log.Fatalln(err.Error())
		// responseHandler.StatusSaveResponse(structs.Jsonresponse{Message: err.Error(), StatusCode: http.StatusBadRequest}, c)
		// return
	}

	c.JSON(200, jsonbody)

}

// list
func (u *UserApi) ListPerson(c *gin.Context) {

	c.JSON(200, nil)
}

// get obj
func (u *UserApi) ViewPerson(c *gin.Context) {

	c.JSON(200, nil)
}
