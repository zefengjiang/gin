package controllers

import (
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {}

func (_ *User) Index(c *gin.Context) {
	userModel := new(models.User)
	list := userModel.List()
	c.HTML(http.StatusOK, "user/index.html", gin.H{"list": list})
}

func (_ *User) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "user/create-edit.html", nil)
}

func (_ *User) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusFound, "/user")
		return
	}

	userModel := new(models.User)
	user := userModel.First(id)
	c.HTML(http.StatusOK, "user/create-edit.html", gin.H{
		"user": user,
	})
}

func (_ *User) Store(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	sex, _ := strconv.Atoi(ctx.PostForm("sex"))
	userModel := new(models.User)
	if id == 0 {
		userModel.Insert(name, age, sex)
	} else {
		userModel.Edit(id, name, age, sex)
	}

	ctx.Redirect(http.StatusFound, "/user")
}

func (_ *User) Del(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Redirect(http.StatusFound, "/user")
		return
	}
	userModel := new(models.User)
	userModel.Del(id)
	ctx.Redirect(http.StatusFound, "/user")
}
