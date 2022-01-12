package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, password, telephone)

	// 判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(newUser)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
