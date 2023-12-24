package libs

import "github.com/jinzhu/gorm"

//全局变量
var DB *gorm.DB


func InitMySQL()(err error){
	dsn := "winner:lmz1995@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	return err
}
