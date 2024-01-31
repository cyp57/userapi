package v1

import (
	ctrlv1 "github.com/cyp57/user-api/app/controller/v1"
	response "github.com/cyp57/user-api/app/api-helper"
)


var authCtrlV1 = new(ctrlv1.AuthCtrl)
var userCtrlV1 = new(ctrlv1.UserCtrl)
var resp   = new(response.ResponseHandlerV2)
