package main

import (
	"Douban/api"
	"Douban/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}