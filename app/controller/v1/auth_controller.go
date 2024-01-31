package v1

import (
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/utils"
)

type AuthCtrl struct{}

func (a *AuthCtrl) Login(data model.LoginInfo) (any,error){

	var appId = utils.GetYaml(cnst.FusionAppId)

	var fusionObj fusionauth.Fusionauth
	fusionObj.LoginId = data.UserName
	fusionObj.Password = data.Password


	fusionObj.SetApplicationId(appId)
	res , err :=fusionObj.Login()
	if err != nil {
		return nil , err
	}

   login := &model.LoginResponse{
	Token:res.Token,
	RefreshToken: res.RefreshToken,
	Uuid: res.User.Id }

return login , nil
}
