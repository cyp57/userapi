package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// response "github.com/cyp57/user-api/app/api-helper"
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
)

type AuthenticationApi struct{}

func (a *AuthenticationApi) Login(c *gin.Context) {
	var loginObj model.LoginInfo
	fmt.Println("loginObj = =", loginObj)
	if err := c.ShouldBindJSON(&loginObj); err != nil {
		resp.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	} else {
		result, err := authCtrlV1.Login(loginObj)
		if err != nil {
			resp.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			resp.SuccessResponse(c, http.StatusOK, result, cnst.LoginSuccess)
			return
		}

	}
}
