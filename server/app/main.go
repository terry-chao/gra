package main

import (
	"gra/api"
	"gra/dao"
)

func main() {
	api.Start()
	dao.DbOperate()
}
