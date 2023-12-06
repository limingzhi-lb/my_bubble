package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//全局变量
var DB *gorm.DB

type ToDO struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
func (todo *ToDO) TableName()string{
	return "todos"
}

func initMySQL()(err error){
	dsn := "winner:lmz1995@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	return err
}

func main(){
	// 连接数据库
	err := initMySQL()
	if err != nil{
		panic(err)
	}
	defer DB.Close()  // 程序退出时关闭数据库
	DB.AutoMigrate(&ToDO{})

	r := gin.Default()
	//去哪里找模版文件引用的静态文件
	r.Static("/static", "static/dist/static/")
	//去哪里加载模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// api v1
	v1 := r.Group("v1")
	v1.POST("/todo", func(ctx *gin.Context) {
		//前端页面填写待办事项，发请求
		var todo ToDO
		ctx.BindJSON(&todo)
		if err := DB.Create(&todo).Error; err != nil{
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, todo)
		}
	})

	//查看所有待办
	v1.GET("/todo", func(ctx *gin.Context) {
		var todoList []ToDO
		if err := DB.Find(&todoList).Error;err!=nil{
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusOK, todoList)
		}
	})

	//查看某一个待办
	v1.GET("/todo/:id", func(ctx *gin.Context) {

	})

	//修改待办
	v1.PUT("/todo/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		var todo ToDO
		res := DB.Where("id=?", id).First(&todo)
		if res.Error != nil{
			ctx.JSON(http.StatusOK, gin.H{"error": res.Error.Error()})
		}
		ctx.BindJSON(&todo)
		if err := DB.Save(&todo).Error; err != nil{
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusOK, todo)
		}
	})

	//删除待办
	v1.DELETE("/todo/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		var todo ToDO
		res := DB.Where("id=?", id).First(&todo)
		if res.Error != nil{
			ctx.JSON(http.StatusOK, gin.H{"error": res.Error.Error()})
		}
		res = DB.Where("id=?", id).Delete(&ToDO{})
		if res.Error != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": res.Error.Error()})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	})

	//添加
	r.Run()
}
