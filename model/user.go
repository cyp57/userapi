package model

import "time"

type RegistrationInfo struct {
	Id        string    `json:"id,omitempty" bson:"id"`
	Uuid      string    `json:"uuid,omitempty" bson:"uuid"`
	Username  string    `json:"username" bson:"username"`
	Email     string    `json:"email" binding:"required" bson:"email"`
	Password  string    `json:"password,omitempty" binding:"required" bson:"password"` // omit from JSON when empty
	FirstName string    `json:"firstName," bson:"firstName"`
	LastName  string    `json:"lastName," bson:"lastName"`
	Age       int       `json:"age," bson:"age"`
	MobilePhone string `json:"mobilePhone," bson:"mobilePhone"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"` // omit from JSON when empty
}

type UserInfo struct {
	Id        string    `json:"id,omitempty"`   // in this case i settle Id to can't edit
	Uuid      string    `json:"uuid,omitempty"` // in this case i settle Uuid to can't edit
	Username  string    `json:"username"`       // in this case i settle Username to can't edit
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password,omitempty"` // edit on fusionauth
	FirstName string    `json:"firstName," bson:"firstName"`
	LastName  string    `json:"lastName," bson:"lastName"`
	Age       int       `json:"age," bson:"age"`
	MobilePhone string `json:"mobilePhone," bson:"mobilePhone"`
	CreatedAt time.Time `json:"created_at,omitempty"` // in this case i settle Username to can't edit
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
