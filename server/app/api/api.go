package api

import (
	"github.com/develop1024/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
	"time"
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

	r.POST("/auth", authHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}

type User struct {
	Username string
	Password string
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("mykey")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

func authHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确

	if checkUser(user.Username, user.Password) {
		// 生成Token
		tokenString, _ := GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

func checkUser(name string, password string) bool {
	db := Before()
	defer db.Close()
	u := User{}
	db.Where("name = ?", "name").First(&u)
	if u.Password == password {
		return true
	}
	return false
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
