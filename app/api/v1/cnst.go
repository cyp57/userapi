package v1

import (
	response "github.com/cyp57/user-api/app/api-helper"
	ctrlv1 "github.com/cyp57/user-api/app/controller/v1"
)

type ApiUserImpl struct {
	IAuthentication
	IUser
}


func InitApiUserImpl() *ApiUserImpl {
	return  &ApiUserImpl{
		IAuthentication: &AuthenticationApi{},
		IUser:            &UserApi{},
	}
}

var authCtrlV1 = new(ctrlv1.AuthCtrl)
var userCtrlV1 = new(ctrlv1.UserCtrl)
var resp = new(response.ResponseHandler)
