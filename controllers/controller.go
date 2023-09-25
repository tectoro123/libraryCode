package controllers

import (
	"BookStore/db"
	"BookStore/models"
	"BookStore/services"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type SearchController struct {
	web.Controller
}

var Data []map[string]interface{}

func (u *SearchController) PostBook() {

	x := getString(u.Ctx.Input.RequestBody)

	services.DbData(x)
	u.Data["json"] = x
	u.ServeJSON()

}
func (u *SearchController) DeleteBook() {
	str := u.Ctx.Input.Param(":id")
	split := strings.Split(str, ",")
	for _, sr := range split {
		a, _ := strconv.Atoi(sr)
		st := make(map[string]interface{})
		db.Connection.Raw("DELETE  FROM `join`.books  WHERE id=?", a).Scan(st)
	}
	u.Data["json"] = "Books Deleted"
	u.ServeJSON()

}

func (u *SearchController) UpdateBook() {
	x := getString(u.Ctx.Input.RequestBody)

	services.DbData(x)
	u.Data["json"] = x
	u.ServeJSON()

}

func (u *SearchController) Search() {
	data := (u.Ctx.Request.Form)
	result := models.Book{}
	var b string
	for key, value := range data {
		b = fmt.Sprintf("%s = '%s'", key, value[0])
	}
	fmt.Printf(b)
	fmt.Println(result)
	db.Connection.Raw("SELECT x.* FROM `join`.books x WHERE ?", b).Scan(&result)
	fmt.Println(result)
	u.Data["json"] = ""
	u.ServeJSON()
}

func (u *SearchController) LibraryBook() {
	st := make([]map[string]interface{}, 0)
	db.Connection.Raw("SELECT x.* FROM `join`.books x").Scan(&st)
	y := make(map[string]interface{})
	x := u.Ctx.Input.Param(":id")
	a, _ := strconv.ParseFloat(x, 64)
	for _, maps := range st {
		fmt.Println(maps["id"].(float64))
		fmt.Println(a)
		if maps["id"].(float64) == a {

			y = maps
			fmt.Println(maps)
		}

	}

	u.Data["json"] = y
	u.ServeJSON()
	fmt.Println(y)

}
func (u *SearchController) GetBook() {
	st := make([]map[string]interface{}, 0)
	db.Connection.Raw("SELECT x.* FROM `join`.books x").Scan(&st)
	u.Data["json"] = st
	u.ServeJSON()

}
func getString(b []byte) []map[string]interface{} {

	var bdata []map[string]interface{}

	json.Unmarshal(b, &bdata)

	return bdata
}
func getMap(b []byte) map[string]interface{} {

	var data map[string]interface{}

	json.Unmarshal(b, &data)

	return data
}
