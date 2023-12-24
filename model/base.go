package model

import (
	"fmt"
	"my_bubble/libs"
)

type BaseModel struct{
	ID int `json:"id"`
}

func (base *BaseModel) Create()bool{
	err := libs.DB.Create(&base).Error
	fmt.Printf("base: %v\n", base)
	if err != nil {
		fmt.Printf("create %v err: %s\n", base, err.Error())
		return false
	}
	return true
}

func (base *BaseModel) Save()bool{
	err := libs.DB.Save(&base).Error
	if err != nil {
		fmt.Printf("save %v err: %s\n", base, err.Error())
		return false
	}
	return true
}

func (base *BaseModel) Delete()bool{
	err := libs.DB.Where("id=?", base.ID).Delete(&base).Error
	if err != nil {
		fmt.Printf("delete %v err: %s\n", base, err.Error())
		return false
	}
	return true
}
