package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/model"
	"github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/pkg/mongodb"
	"github.com/cyp57/user-api/setting"
	"github.com/cyp57/user-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// edit all value
func (u *UserCtrl) EditUser(uuid string, data *model.UserInfo) (interface{}, error) {

	var fusionObj fusionauth.Fusionauth
	fusionObj.Email = data.Email
	fusionObj.FirstName = data.FirstName
	fusionObj.LastName = data.LastName
	fusionObj.MobilePhone = data.MobilePhone

	_, err := fusionObj.PatchUser(uuid)
	if err != nil {
		return nil, err
	}

	now, err := utils.GetCurrentTime()
	if err != nil {
		return nil, err
	}
	data.UpdatedAt = now
	m, err := utils.StructToM(data)
	if err != nil {
		return nil, err
	}
	utils.Debug(m)

	collectionName := setting.CollectionSetting.User
	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": m}
	_, err = mongodb.UpdateDocument(collectionName, filter, update, nil)
	if err != nil {
		return nil, err
	}

	return uuid, nil
}

// edit only the values ​​sent.
func (u *UserCtrl) PatchUser(uuid string, data map[string]interface{}) (interface{}, error) {

	var fusionObj fusionauth.Fusionauth
	if data["email"] != nil {
		fusionObj.Email = fmt.Sprint(data["email"])
	}
	if data["firstName"] != nil {
		fusionObj.FirstName = fmt.Sprint(data["firstName"])
	}
	if data["lastName"] != nil {
		fusionObj.LastName = fmt.Sprint(data["lastName"])
	}
	if data["mobilePhone"] != nil {
		fusionObj.MobilePhone = fmt.Sprint(data["mobilePhone"])
	}

	_, err := fusionObj.PatchUser(uuid)
	if err != nil {
		return nil, err
	}

	/// in this case i settle for 4 field that can edit(mobilePhone,lastName,firstName,email)
	newMap := make(map[string]interface{})
	if data["email"] != nil {
		newMap["email"] = data["email"]
	}
	if data["firstName"] != nil {
		newMap["firstName"] = data["firstName"]
	}
	if data["lastName"] != nil {
		newMap["lastName"] = data["lastName"]
	}
	if data["mobilePhone"] != nil {
		newMap["mobilePhone"] = data["mobilePhone"]
	}

	now, err := utils.GetCurrentTime()
	if err != nil {
		return nil, err
	}
	newMap["updated_at"] = now

	collectionName := setting.CollectionSetting.User
	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": newMap}
	utils.Debug(update)
	_, err = mongodb.UpdateDocument(collectionName, filter, update, nil)
	if err != nil {
		return nil, err
	}

	return uuid, err
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

func (u *UserCtrl) GetUserList(filter map[string]interface{}) (result *[]model.UserInfo, err error) {
	userList := make([]model.UserInfo, 0)

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	pipeline := []bson.M{}

	if len(filter) > 0 {
		if filter["search"] != nil {
			txtStr := fmt.Sprint(filter["search"])
			search := bson.M{
				"$or": []bson.M{
					{
						"firstName": bson.M{
							"$regex": primitive.Regex{
								Pattern: txtStr,
								Options: "ix",
							},
						},
					},
					{
						"lastName": bson.M{
							"$regex": primitive.Regex{
								Pattern: txtStr,
								Options: "ix",
							},
						},
					},
					{
						"email": bson.M{
							"$regex": primitive.Regex{
								Pattern: txtStr,
								Options: "ix",
							},
						},
					},
				},
			}
			pipeline = append(pipeline, bson.M{"$match": search})
		}

		if filter["uuid"] != nil {
			pipeline = append(pipeline, bson.M{"$match": bson.M{"uuid": filter["uuid"]}})
		}

		if filter["sort"] != nil && filter["sortkey"] != nil {
			sort, err := strconv.Atoi(fmt.Sprint(filter["sort"]))
			if err != nil {
				panic(err)
			}

			keyName := fmt.Sprint(filter["sortkey"])
			pipeline = append(pipeline, bson.M{"$sort": bson.M{keyName: sort}})

		} else {
			if filter["sort"] != nil { // default : update_at
				sort, err := strconv.Atoi(fmt.Sprint(filter["sort"]))
				if err != nil {
					panic(err)
				}
				pipeline = append(pipeline, bson.M{"$sort": bson.M{"updated_at": sort}})
			} else if filter["sortkey"] != nil {
				return nil, errors.New(cnst.ErrSortKeyReq)
			}
		}

		if filter["page"] != nil {
			if filter["limit"] != nil {
				limit, err := strconv.Atoi(fmt.Sprint(filter["limit"]))
				if err != nil {

					panic(err)
				}
				offset, err := strconv.Atoi(fmt.Sprint(filter["page"]))
				if err != nil {

					panic(err)
				}

				skip := limit * (offset - 1)
				pipeline = append(pipeline, bson.M{"$skip": skip})
			}
		}

		if filter["limit"] != nil {
			limit, err := strconv.Atoi(fmt.Sprint(filter["limit"]))
			if err != nil {

				panic(err)
			}
			pipeline = append(pipeline, bson.M{"$limit": limit})
		}

		fmt.Println("pipeline show")
		utils.Debug(pipeline)
	}

	collectionName := setting.CollectionSetting.User
	userData, err := mongodb.AggregateDocument(collectionName, pipeline)
	if err != nil {
		return nil, err
	}

	if len(userData) > 0 {
		bytes, err := json.Marshal(&userData)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bytes, &userList)
		if err != nil {
			return nil, err
		}
	}

	return &userList, err
}

func (u *UserCtrl) DeleteUser(uuid string) (interface{}, error) {
	errCh := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := new(fusionauth.Fusionauth).DeleteUser(uuid)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		filter := bson.M{"uuid": uuid}
		collectionName := setting.CollectionSetting.User
		_, err := mongodb.DeleteOneDocument(collectionName, filter)
		if err != nil {
			errCh <- err
			return
		}
	}()

	wg.Wait()
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return uuid, nil
}
