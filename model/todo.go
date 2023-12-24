package model

import (
	"fmt"
	"my_bubble/libs"
)

type ToDO struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func (todo *ToDO) TableName()string{
	return "todos"
}

func (todo ToDO) Create()bool{
	err := libs.DB.Create(&todo).Error
	if err != nil {
		fmt.Println("create todo err, ", err.Error())
		return false
	}
	return true
}

func (todo ToDO) Save()bool{
	err := libs.DB.Save(&todo).Error
	if err != nil {
		fmt.Println("save todo err, ", err.Error())
		return false
	}
	return true
}

func (todo ToDO) Delete()bool{
	err := libs.DB.Where("id=?", todo.ID).Delete(&todo).Error
	if err != nil {
		fmt.Println("delete todo err, ", err.Error())
		return false
	}
	return true
}

func (todo ToDO) Get(id string) (bool, ToDO){

	res := libs.DB.Where("id=?", id).First(&todo)
	if res.Error != nil{
		fmt.Println("get todo err, ", res.Error.Error())
		return false, todo
	}
	return true, todo
}

// todo 结构体的方法，怎么区分类方法和实例方法？