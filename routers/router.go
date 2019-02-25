package routers

import (
	"gin/controllers"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	user := new(controllers.User)
	v1 := r.Group("/")
	{
		v1.GET("/user", user.Index)
		v1.GET("/user/create", user.Create)
		v1.GET("/user/:id", user.Edit)
		v1.DELETE("/article/:id", user.Del)
		v1.POST("/article/store", user.Store)
	}

	return r
}
