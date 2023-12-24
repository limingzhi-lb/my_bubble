package model

import (
	"time"
)

type User struct{
	ID int `json:"id"`
	Name int `json:"name"`
	Password string `json:"password"`
	UpdateTime time.Time  `json:"update_time"`
	CreateTime time.Time  `json:"create_time"`
}


func (user *User) TableName()string{
	return "users"
}

