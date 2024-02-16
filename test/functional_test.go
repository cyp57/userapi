package test

import (
	"fmt"
	"testing"

	ctrlv1 "github.com/cyp57/user-api/app/controller/v1"
	"github.com/cyp57/user-api/cnst"
	"github.com/cyp57/user-api/config"
	"github.com/cyp57/user-api/model"
	fusionauthPkg "github.com/cyp57/user-api/pkg/fusionauth"
	"github.com/cyp57/user-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/cyp57/user-api/pkg/mongodb"
	"github.com/cyp57/user-api/setting"
)

const (
	pathEnv  = "../config/.env"
	pathIni  = "../config/app.ini"
	pathYaml = "../config"
)

var ctrlV1 = iniCtrlV1()

type testCtrlV1 struct {
	ctrlv1.IAuthCtrl
	ctrlv1.IUserCtrl
}

func iniCtrlV1() *testCtrlV1 {
	return &testCtrlV1{
		IAuthCtrl: &ctrlv1.AuthCtrl{},
		IUserCtrl: &ctrlv1.UserCtrl{},
	}
}

func initTestEnvironment() {
	config.LoadConfig(pathEnv, pathYaml)
	setting.InitIni(pathIni)
	mongodb.MongoDbConnect()
	fusionauthPkg.InitFusionAuth()
}

func Test_CreateUser(t *testing.T) {
	initTestEnvironment()
	type userData struct {
		userName string
		password string
		email    string
	}

	type testCreateUser struct {
		data   userData
		isErr  bool
	}

	tests := []testCreateUser{
		{
			data: userData{
				userName: "user1",
				password: "12345678",
				email:    "user1@example.com",
			},
			isErr: false,
		},
		{
			data: userData{
				userName: "user2",
				password: "12345678",
				email:    "user1@example.com",
			},
			isErr: true,
		},
	}

	var resultUuid = make([]string, 0)

	fusionAppId := utils.GetYaml(cnst.FusionAppId)
	for _, test := range tests {
		user := &model.RegistrationInfo{
			Username: test.data.userName,
			Email:    test.data.email,
			Password: test.data.password,
		}

		result, err := ctrlV1.CreateUser(user, fusionAppId)

		if !test.isErr && err != nil {
			t.Errorf("Expected no error, but got: %v from test data: %v", err.Error(), *user)
			return
		}
		if test.isErr && err == nil {
			t.Errorf("Expected an error, but got nil from test data: %v", *user)
			return
		}

		if result != nil {
			resultUuid = append(resultUuid, fmt.Sprint(result.(primitive.M)["uuid"]))
		}
	}
	for _, uuid := range resultUuid {
		_, err := ctrlV1.GetUserInfo(uuid, nil)
		if err != nil {
			t.Errorf("CreateUser failed. Expected user data with uuid : %v, but got error :%v", uuid, err.Error())
		} else { // delete user
			_, err := ctrlV1.DeleteUser(uuid)
			if err != nil {
				t.Errorf("DeleteUser failed. Expected delete user after create with uuid : %v, but got error :%v", uuid, err.Error())
			}
		}
	}

}

func Test_GetUserInfo(t *testing.T) {

	type testGetUserInfo struct {
		id     string
		isErr  bool
		expect interface{}
	}

	initTestEnvironment()

	tests := []testGetUserInfo{
		{
			id:     "99997719-333d-4e41-b22a-6432be2df115", // not available in database
			isErr:  true,
			expect: "mongo: no documents in result", // get data with the expected error message
		},
		{
			id:     "03597719-333d-4e41-b22a-6432be2df115", // available in database
			isErr:  false,
			expect: "03597719-333d-4e41-b22a-6432be2df115",
		},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			result, err := ctrlV1.GetUserInfo(test.id, nil)

			if test.isErr && err == nil {
				t.Errorf("Expected an error, but got nil for test ID: %v", test.id)
				return
			} else if !test.isErr && err != nil {
				t.Errorf("Expected no error, but got: %v for test ID: %v", err.Error(), test.id)
				return
			}

			if !test.isErr && err == nil && result.Uuid != test.expect {
				t.Errorf("Expected userInfo with uuid %v, but got %v for test ID: %v", test.expect, result.Uuid, test.id)
			} else if test.isErr && err != nil && err.Error() != test.expect {
				t.Errorf("Expected error message %v, but got %v for test ID: %v", test.expect, err.Error(), test.id)
			}
		})
	}

}
