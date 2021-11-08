package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"gra/auth"
	"gra/viper"
	_ "gra/viper"
	"strconv"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

var username string
var password string
var address string
var tableName string

var mysqlUrl = username + ":" + password + "@(" + address + ")/" + tableName + "?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	config := viper.GetMysqlConfig()
	username = config.MysqlInfo.Username
	password = config.MysqlInfo.Password
	address = config.MysqlInfo.Address
	tableName = config.MysqlInfo.Table_name
}

func DbInit() {
	viper.GetMysqlConfig()
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&UserInfo{})
}

func Before() (db *gorm.DB) {
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}

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
	r.GET("/id", func(c *gin.Context) {
		idString := c.Query("id")
		id, _ := strconv.Atoi(idString)
		user := getUser(id)
		c.JSON(200, gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"gender": user.Gender,
			"hobby":  user.Hobby,
		})

	})

	//r.POST("/auth", auth.AuthHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func addUser(id int, name string, gender string, hobby string) {
	db := Before()
	defer db.Close()
	// 创建记录
	u := UserInfo{ID: uint(id), Name: name, Gender: gender, Hobby: hobby}
	db.Create(&u)
}

func deleteUser(id int, name string, gender string, hobby string) {
	db := Before()
	defer db.Close()
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
	defer db.Close()
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
	defer db.Close()
	// 查询
	var u = new(UserInfo)
	db.First(&u, id)
	return u
}
