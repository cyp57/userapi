package model

type LoginInfo struct {

	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`

}


type LoginResponse struct {
 
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Uuid string `json:"uuid"`
	
}