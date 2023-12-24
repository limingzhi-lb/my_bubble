package main

import (
	apiV1 "my_bubble/api/v1"
	"my_bubble/libs"
	"my_bubble/views"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
	// 连接数据库
	err := libs.InitMySQL()
	if err != nil{
		panic(err)
	}
	defer libs.DB.Close()  // 程序退出时关闭数据库

	r := gin.Default()
	//去哪里找模版文件引用的静态文件
	r.Static("/static", "static/dist/static/")
	//去哪里加载模版文件
	r.LoadHTMLGlob("templates/*")
	views.BoundViews(r)

	// api v1
	v1 := r.Group("v1")
	apiV1.ToDoApi(v1)

	r.Run()
}
