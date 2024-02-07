package v1

import (
	"encoding/json"
	"fmt"

	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/pkg/mongodb"
	"github.com/cyp57/user-api/setting"
	"github.com/cyp57/user-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type UserCtrl struct{}

func (u *UserCtrl) CreateUser(data *model.RegistrationInfo, appId string, isAdmin ...bool) (interface{}, error) {

	var fusionObj fusionauth.Fusionauth
	if len(isAdmin) > 0 {
		if isAdmin[0] {
			fusionObj.Roles = []string{"admin"}
		}
	} else { // default customer
		fusionObj.Roles = []string{"customer"}
	}
	fusionObj.Username = data.Username
	fusionObj.Email = data.Email
	fusionObj.Password = data.Password
	fusionObj.FirstName = data.FirstName
	fusionObj.LastName = data.LastName
	fusionObj.SetApplicationId(appId)
	// fusionObj.MobilePhone = data.

	resp, err := fusionObj.Register()
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil, err
	}
	fmt.Println("resp.Token :", resp.Token)
	fmt.Println("resp.RefreshToken :", resp.RefreshToken)

	data.Uuid = resp.User.Id
	hashed, err := utils.HashPassword(data.Password)
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil, err
	}
	data.Password = hashed
	now, err := utils.GetCurrentTime()
	if err != nil {
		return nil, err
	}
	data.CreatedAt = now
	data.UpdatedAt = now

	m, err := utils.StructToM(data)
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil, err
	}
	collectionName := setting.CollectionSetting.User
	fmt.Println("collectionName = = ", collectionName)
	result, err := mongodb.InsertOneDocument(collectionName, m, "Uc")
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil, err
	}

	return result, nil
}

func (u *UserCtrl) EditUser(uuid string, data *model.UserInfo) (interface{}, error) {

	var fusionObj fusionauth.Fusionauth
	data.Email = data.Email

	// fusionObj.Email = data.Email
	// fusionObj.FirstName = data.FirstName
	// fusionObj.LastName = data.LastName
	_, err := fusionObj.PatchUser(uuid)
	// if err != nil {
	// 	fmt.Println("err :", err.Error())
	// 	return nil, err
	// }






	now, err := utils.GetCurrentTime()
	if err != nil {
		return nil, err
	}
	data.UpdatedAt = now
	m, err := utils.StructToM(data)
	if err != nil {
		fmt.Println("err :", err.Error())
		return nil, err
	}

	fmt.Println("StructToM = = ", m)
	utils.Debug(m)
	return nil, nil
}

func (u *UserCtrl) GetUserInfo(uuid string) (*model.UserInfo, error) {

	collectionName := setting.CollectionSetting.User
	filter := bson.M{"uuid": uuid}
	userData, err := mongodb.FindOneDocument(collectionName, filter)
	utils.Debug(userData)
	if err != nil {
		return nil, err
	}

	userObj := &model.UserInfo{}
	bytes, err := json.Marshal(&userData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, userObj)
	if err != nil {
		return nil, err
	}

	return userObj, nil
}
