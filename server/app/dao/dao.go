package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func DbOperate() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1)/gra?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{ID: 1, Name: "terry", Gender: "男", Hobby: "篮球"}

	// 创建记录
	db.Create(&u1)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)

	db.AutoMigrate(&UserInfo{})

}
