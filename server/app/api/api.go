package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func DbInit() {
	mysqlUrl := "root:root@(127.0.0.1)/gra?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&UserInfo{})
}

func Before() (db *gorm.DB) {
	mysqlUrl := "root:root@(127.0.0.1)/gra?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func Start() {
	r := gin.Default()

	// add
	r.POST("/add_user", func(c *gin.Context) {
		idString := c.PostForm("id")
		id, _ := strconv.Atoi(idString)
		name := c.PostForm("name")
		gender := c.PostForm("gender")
		hobby := c.PostForm("hobby")
		addUser(id, name, gender, hobby)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// delete
	r.DELETE("/id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		// 删除
	})
	// update
	r.PUT("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// select
	r.GET("/id/:id/", func(c *gin.Context) {
		idString := c.Param("id")
		id, _ := strconv.Atoi(idString)
		getUser(id)
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func addUser(id int, name string, gender string, hobby string) {
	db := Before()
	// 创建记录
	u := UserInfo{ID: uint(id), Name: name, Gender: gender, Hobby: hobby}
	db.Create(&u)
}

func deleteUser(id int, name string, gender string, hobby string) {
	db := Before()
	// 删除
	u := &UserInfo{
		ID:     uint(id),
		Name:   name,
		Gender: gender,
		Hobby:  hobby,
	}
	db.Delete(&u)
	db.AutoMigrate(&UserInfo{})
}

func putUser(id int, name string, gender string, hobby string) {
	db := Before()
	// 更新
	u := &UserInfo{
		ID:     uint(id),
		Name:   name,
		Gender: gender,
		Hobby:  hobby,
	}
	db.Model(&u).Update("hobby", "双色球")
}

func getUser(id int) *UserInfo {
	db := Before()

	// 查询
	var u = new(UserInfo)
	db.First(u)
	db.First(&u, id)
	db.AutoMigrate(&UserInfo{})
	return &UserInfo{}
}
