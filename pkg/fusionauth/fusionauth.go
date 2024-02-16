package fusionauth

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/sirupsen/logrus"
	lrlog "github.com/cyp57/user-api/pkg/logrus"
	resHandler "github.com/cyp57/user-api/app/api-helper"
	"github.com/cyp57/user-api/cnst"
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

func InitFusionAuth() {
	var host = utils.GetYaml(cnst.FusionHost)
	var apiKey = utils.GetYaml(cnst.FusionAPIKey)
	var httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	var baseURL, err = url.Parse(host)
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}
	// Construct a new FusionAuth Client
	Auth := fusionauth.NewClient(httpClient, baseURL, apiKey)

	// for production code, don't ignore the error!
	_, err = Auth.RetrieveTenants()
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}

	fusionauthConnection(Auth)
}

func fusionauthConnection(client *fusionauth.FusionAuthClient) {
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

	response, restErr, err := AuthClient.Login(request)
	utils.Debug("AuthClient.Login :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4000)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4001)
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


	response, restErr, err := AuthClient.Register("", request)
	utils.Debug("AuthClient.Register :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4002)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4003)
		return nil, restErr
	}

	return response, nil
}

func (f *Fusionauth) PatchUser(uuid string) (*fusionauth.UserResponse, error) {

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
	utils.Debug("AuthClient.PatchUser :")
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4004)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4005)
		return nil, restErr
	}

	return response, nil
}

func (f *Fusionauth) PatchRegistration(uuid string) (*fusionauth.RegistrationResponse, error) {

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
	utils.Debug("AuthClient.PatchRegistration :")
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4006)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4007)
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
	utils.Debug("AuthClient.ForgotPassword :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4008)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4009)
		return nil, restErr
	}

	return response, nil
}
func (f *Fusionauth) ChangePassword(currentPassword string, newPassword string) error {
	var request fusionauth.ChangePasswordRequest
	request.LoginId = f.LoginId
	request.CurrentPassword = currentPassword
	request.Password = newPassword
	request.ApplicationId = f.ApplicationId

	response, restErr, err := AuthClient.ChangePasswordByIdentity(request)
	utils.Debug("AuthClient.ChangePasswordByIdentity :")
	utils.Debug(response)
	if err != nil {
		resHandler.SetErrorCode(4010)
		return err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4011)
		return restErr
	}

	return nil
}

func (f *Fusionauth) ValidateToken(token string) (*fusionauth.ValidateResponse, error) {

	response, err := AuthClient.ValidateJWT(token)
	utils.Debug("AuthClient.ValidateJWT :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4012)
		return nil, err
	}

	return response, nil
}

func (f *Fusionauth) GetUserRegistration(uuid string) (*fusionauth.RegistrationResponse, error) {

	if utils.IsEmptyString(f.ApplicationId) {
		return nil, fmt.Errorf(cnst.ErrappIdNotFound)
	}

	response, restErr, err := AuthClient.RetrieveRegistration(uuid, f.ApplicationId)
	utils.Debug("AuthClient.RetrieveRegistration :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4013)
		return nil, err
	}

	if restErr != nil {
		resHandler.SetErrorCode(4014)
		return nil, restErr
	}

	return response, nil
}

func (f *Fusionauth) DeleteUser(uuid string) error {
	_, restErr, err := AuthClient.DeleteUser(uuid)
	if err != nil {
		resHandler.SetErrorCode(4015)
		return err
	}
	if restErr != nil {
		resHandler.SetErrorCode(4016)
		return restErr
	}

	return nil
}

func (f *Fusionauth) NewAccessToken(token, refreshToken string) (*fusionauth.IssueResponse, error) {

	response, restErr, err := AuthClient.IssueJWT(f.ApplicationId, token, refreshToken)
	utils.Debug("AuthClient.IssueJWT :")
	utils.Debug(response)

	if err != nil {
		resHandler.SetErrorCode(4017)
		return nil, err
	}
	if restErr != nil {
		resHandler.SetErrorCode(4018)
		return nil, restErr
	}

	return response, nil
}
