package v1

import (
	

	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/utils"
)


type IAuthCtrl interface{
	Login(*model.LoginInfo, string) (interface{}, error)
	RefreshJwt(*model.RefreshJwt, string) (interface{}, error)
}

type AuthCtrl struct{}

func (a *AuthCtrl) Login(data *model.LoginInfo, appId string) (interface{}, error) {

	var fusionObj fusionauth.Fusionauth
	fusionObj.LoginId = data.UserName
	fusionObj.Password = data.Password

	fusionObj.SetApplicationId(appId)
	res, err := fusionObj.Login()
	if err != nil {
		return nil, err
	}

	resReg, errReg := fusionObj.GetUserRegistration(res.User.Id)
	utils.Debug("GetUserRegistration :")
	utils.Debug(resReg)

	if errReg != nil {
		return nil, err
	}

	login := &model.LoginResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
		Uuid:         res.User.Id,
		Roles:        resReg.Registration.Roles}

	return login, nil
}

func (a *AuthCtrl) RefreshJwt(data *model.RefreshJwt, appId string) (interface{}, error) {


	var fusionObj fusionauth.Fusionauth
	fusionObj.SetApplicationId(appId)

	res ,err := fusionObj.NewAccessToken(data.Token,data.RefreshToken)
	if err != nil {
		return nil, err
	}
	utils.Debug("NewAccessToken ")
	utils.Debug(res)

	return res, nil
}