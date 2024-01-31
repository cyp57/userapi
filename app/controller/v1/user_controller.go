package v1

import (
	"fmt"
	

	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/pkg/mongodb"
	"github.com/cyp57/user-api/setting"
	"github.com/cyp57/user-api/utils"
)

type UserCtrl struct{}

func (u *UserCtrl) CreateUser(data *model.UserInfo, appId string) (any,error){

	var fusionObj fusionauth.Fusionauth

	fusionObj.Username = data.Username
	fusionObj.Email = data.Email
	fusionObj.Password = data.Password
	fusionObj.FirstName = data.FirstName
	fusionObj.LastName = data.LastName
	// fusionObj.MobilePhone = data.

	resp, err := fusionObj.Register(appId)
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil , err
	}
	fmt.Println("resp.Token :", resp.Token)
	fmt.Println("resp.RefreshToken :", resp.RefreshToken)


	data.Uuid = resp.User.Id
	hashed,err:=utils.HashPassword(data.Password)
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil , err
	}
	data.Password = hashed
	now ,err:=utils.GetCurrentTime()
	if err != nil {
		return nil , err
	}
	data.CreatedAt = now
	data.UpdatedAt = now

	m,err:= utils.StructToM(data)
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil , err
	}
	collectionName := setting.CollectionSetting.User
	fmt.Println("collectionName = = ",collectionName)
	result,err:= mongodb.InsertOneDocument(collectionName,m,"Uc")
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil , err
	}


	return result , nil
}
