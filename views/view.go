package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func BoundViews(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
