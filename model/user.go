package model

import "time"

type User struct {
	Id          string `json:"id,omitempty"`
	UserName    string `json:"userName"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"` // encode when save in db
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MobilePhone string `json:"mobilePhone,"`
	Uuid        string `json:"uuid"` //for fusionauth

	// `json:"username" "binding:"required"`    . .. . . `json:"uuid,omitempty"`
}

type UserInfo struct {
	Id        string    `json:"id,omitempty" bson:"id"`
	Uuid       string  `json:"uuid,omitempty" bson:"uuid"`

	Username  string    `json:"username" bson:"username"`
	Email     string    `json:"email" binding:"required" bson:"email"`
	Password  string    `json:"password,omitempty" binding:"required" bson:"password"` // omit from JSON when empty
	FirstName string    `json:"firstName," bson:"firstName"`
	LastName  string    `json:"lastName," bson:"lastName"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"` // omit from JSON when empty
}
