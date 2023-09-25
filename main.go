package main

import (
	"BookStore/db"
	_ "BookStore/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	db.ConnectToDB()
	web.Run()
}
