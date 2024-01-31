package fusionauth

import (
	"encoding/json"
	"fmt"
	"github.com/FusionAuth/go-client/pkg/fusionauth"

	resHandler "github.com/cyp57/user-api/app/api-helper"
)

type IFusionauth interface{
	InitConnection(*fusionauth.FusionAuthClient) 
	Login() (*fusionauth.LoginResponse, error)
}


type (
	Fusionauth struct{
		ApplicationId string
		LoginId string // username or email
		Password string

		Username string
		Email  string
		FirstName  string
		LastName  string
		MobilePhone  string

	}



)

var AuthClient *fusionauth.FusionAuthClient

func (f *Fusionauth) InitConnection(client *fusionauth.FusionAuthClient) {
	fmt.Println("InitConnection= = ", client)
	AuthClient = client
}

func (f *Fusionauth) SetApplicationId(appId string) {
	f.ApplicationId = appId
}



func (f *Fusionauth) Login() (response *fusionauth.LoginResponse , err error){
	var request fusionauth.LoginRequest
	request.ApplicationId = f.ApplicationId
	request.LoginId = f.LoginId
	request.Password = f.Password
    
	response, restErr , err := AuthClient.Login(request) // 
	fmt.Println("response =-= ",response)
	fmt.Println("restErr =-= ",restErr)
	fmt.Println("err =-= ",err)
if err != nil {
	resHandler.SetErrorCode(4003)
	return nil , err
}

if restErr != nil {
	resHandler.SetErrorCode(4004)
	return nil , restErr
}


	if response != nil  && err == nil {
		jsMar,err:= json.Marshal(response)
		if err != nil {
		fmt.Println("Marshal err :", err.Error())	
		}
		fmt.Println("json : ",string(jsMar)) 
		
	} else { // connection fusionauth failed
		fmt.Println("connection fusionauth failed" )
	}

	return response , nil
}

func (f *Fusionauth) Register(appId string) (response *fusionauth.RegistrationResponse , err error){ 
	var request fusionauth.RegistrationRequest
	request.Registration.ApplicationId = appId
	request.GenerateAuthenticationToken = true
	request.User.Username = f.Username
	request.User.Email = f.Email
	request.User.Password = f.Password
	request.User.FirstName = f.FirstName
	request.User.LastName = f.LastName
	request.User.MobilePhone = f.MobilePhone

	response , restErr , err := AuthClient.Register("",request)
	fmt.Println("AuthClient.Register response : ",response)
	fmt.Println("AuthClient.Register err : ",err)
	fmt.Println("AuthClient.Register restErr : ",restErr)
	if err != nil {
		resHandler.SetErrorCode(4001)
		return nil , err
	}

	if restErr != nil {
	    resHandler.SetErrorCode(4002)
		return nil , restErr
	}


	if response != nil { 
	fmt.Println("err")
	by , err := json.Marshal(response)
	 if err != nil {
		fmt.Println("marshal : ",err.Error())
	 }
	 fmt.Println("marshal json :",string(by))
	}

	 return response , nil
}

func (f *Fusionauth) ForgotPassword(){}
func (f *Fusionauth) ChangePassword(){}

func (f *Fusionauth) ValidateToken(){}
