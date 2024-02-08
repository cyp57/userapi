package fusionauth

import (
	"fmt"

	"github.com/FusionAuth/go-client/pkg/fusionauth"

	resHandler "github.com/cyp57/user-api/app/api-helper"
	"github.com/cyp57/user-api/utils"
)


type (
	Fusionauth struct {
		ApplicationId string
		LoginId       string // username or email
		Password      string

		Username    string
		Email       string
		FirstName   string
		LastName    string
		MobilePhone string
		Roles       []string
	}
)

var AuthClient *fusionauth.FusionAuthClient

func (f *Fusionauth) InitConnection(client *fusionauth.FusionAuthClient) {
	AuthClient = client
}

func (f *Fusionauth) SetApplicationId(appId string) {
	f.ApplicationId = appId
}

func (f *Fusionauth) Login() (response *fusionauth.LoginResponse, err error) {
	var request fusionauth.LoginRequest
	request.ApplicationId = f.ApplicationId
	request.LoginId = f.LoginId
	request.Password = f.Password

	response, restErr, err := AuthClient.Login(request) //
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4003)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4004)
		return nil, restErr
	}

	return response, nil
}

func (f *Fusionauth) Register() (response *fusionauth.RegistrationResponse, err error) {
	var request fusionauth.RegistrationRequest
	request.Registration.ApplicationId = f.ApplicationId

	request.GenerateAuthenticationToken = true
	request.User.Username = f.Username
	request.User.Email = f.Email
	request.User.Password = f.Password
	request.User.FirstName = f.FirstName
	request.User.LastName = f.LastName
	request.User.MobilePhone = f.MobilePhone
	request.Registration.Roles = f.Roles
	
	fmt.Println("func : Register : ", request)
	response, restErr, err := AuthClient.Register("", request)
	utils.Debug(response)
	fmt.Println("AuthClient.Register err : ", err)
	fmt.Println("AuthClient.Register restErr : ", restErr)
	if err != nil {
		resHandler.SetErrorCode(4001)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4002)
		return nil, restErr
	}

	return response, nil
}



func (f *Fusionauth) PatchUser(uuid string) (*fusionauth.UserResponse, error) {

	// request := map[string]interface{}{
	// 	"user": map[string]interface{}{
	// 		"firstName":   f.FirstName,
	// 		"lastName":    f.LastName,
	// 		"email":       f.Email,
	// 		"mobilePhone": f.MobilePhone,
	// 	},
	// }

	request := make(map[string]interface{})
	user := make(map[string]interface{})
	if !utils.IsEmptyString(f.FirstName) {
		user["firstName"] = f.FirstName
	}
	if !utils.IsEmptyString(f.LastName) {
		user["lastName"] = f.LastName
	}
	if !utils.IsEmptyString(f.Email) {
		user["email"] = f.Email
	}
	if !utils.IsEmptyString(f.MobilePhone) {
		user["mobilePhone"] = f.MobilePhone
	}
	request["user"] = user

	response, restErr, err := AuthClient.PatchUser(uuid, request)
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4001)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4002)
		return nil, restErr
	}
	
	return response, nil
}

func (f *Fusionauth) PatchRegistration(uuid string) (*fusionauth.RegistrationResponse, error) {
	// var request fusionauth.UserRequest

	// request.User.Username = f.Username
	// request.User.Email = f.Email
	// request.User.FirstName = f.FirstName
	// request.User.LastName = f.LastName
	// request.User.MobilePhone = f.MobilePhone

	// Example request to update user fields
	// request := map[string]interface{}{
	// 	"registration": map[string]interface{}{
	// 		"applicationId" : "",
	// 		"username": f.Username, // Update the first name
	// 		"roles":  f.Roles,  // Update the last name
	// 	},

	// 	// Add more fields as needed
	// }

	request := make(map[string]interface{})
	registration := make(map[string]interface{})
	if !utils.IsEmptyString(f.ApplicationId) {
		registration["applicationId"] = f.ApplicationId
	}
	if !utils.IsEmptyString(f.Username) {
		registration["username"] = f.Username
	}
	if len(f.Roles) > 0 {
		registration["roles"] = f.Roles
	}
	request["registration"] = registration

	response, restErr, err := AuthClient.PatchRegistration(uuid, request)
	// response,restErr,err:=AuthClient.UpdateUser(uuid,request)
	fmt.Println("AuthClient.UpdateUser :")
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4005)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4006)
		return nil, restErr
	}

	return response, nil
}

func (f *Fusionauth) ForgotPassword() (*fusionauth.ForgotPasswordResponse, error) {
	var request fusionauth.ForgotPasswordRequest
	request.ApplicationId = f.ApplicationId
	request.LoginId = f.LoginId
	request.SendForgotPasswordEmail = true

	response, restErr, err := AuthClient.ForgotPassword(request)
	fmt.Println("AuthClient.ForgotPassword :")
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4007)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4008)
		return nil, restErr
	}

	return response, nil
}
func (f *Fusionauth) ChangePassword(currentPassword string, newPassword string) error {
	var request fusionauth.ChangePasswordRequest
	request.LoginId = f.LoginId
	request.CurrentPassword = currentPassword
	request.Password = newPassword

	response, restErr, err := AuthClient.ChangePasswordByIdentity(request)
	fmt.Println("AuthClient.ChangePassword :")
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4009)
		return err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4010)
		return restErr
	}

	return nil
}

func (f *Fusionauth) ValidateToken(token string) (*fusionauth.ValidateResponse, error){

	response, err := AuthClient.ValidateJWT(token)
	fmt.Println("AuthClient.ValidateJWT :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4011)
		return nil , err
	}

 return response , nil
}

func (f *Fusionauth) GetUserRegistration(uuid string) (*fusionauth.RegistrationResponse, error){

	if utils.IsEmptyString(f.ApplicationId ) {
		return nil , fmt.Errorf("fusionauth applicationId not found")
	}

	response ,restErr,err:=AuthClient.RetrieveRegistration(uuid , f.ApplicationId)
	if err != nil {
		resHandler.SetErrorCode(4009)
		return nil , err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4010)
		return nil , restErr
	}

	return response , nil
}
