package v1

import (
	"my_bubble/libs"
	"my_bubble/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ToDoApi(v1 *gin.RouterGroup){
	v1.POST("/todo", func(ctx *gin.Context) {
		//前端页面填写待办事项，发请求
		var todo model.ToDO
		ctx.BindJSON(&todo)
		isOk := todo.Create()
		if !isOk{
			ctx.JSON(http.StatusOK, gin.H{
				"error": "create failed",
			})
		} else {
			ctx.JSON(http.StatusOK, todo)
		}
	})

	//查看所有待办
	v1.GET("/todo", func(ctx *gin.Context) {
		var todoList []model.ToDO
		if err := libs.DB.Find(&todoList).Error;err!=nil{
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
		var todo model.ToDO
		res := libs.DB.Where("id=?", id).First(&todo)
		if res.Error != nil{
			ctx.JSON(http.StatusOK, gin.H{"error": res.Error.Error()})
		}
		ctx.BindJSON(&todo)
		isOk := todo.Save()
		if !isOk{
			ctx.JSON(http.StatusOK, gin.H{"error": "save failed"})
		}else{
			ctx.JSON(http.StatusOK, todo)
		}
	})

	//删除待办
	v1.DELETE("/todo/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		isOk, todo := model.ToDO{}.Get(id)
		isOk = model.ToDO{}.Delete()

		if !isOk{
			ctx.JSON(http.StatusOK, gin.H{"error": "query failed"})
		}

		isOk = todo.Delete()
		if !isOk {
			ctx.JSON(http.StatusOK, gin.H{"error": "delete failed"})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	})

}