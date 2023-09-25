package routers

import (
	"BookStore/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/books", &controllers.SearchController{}, "post:PostBook")
	web.Router("/books", &controllers.SearchController{}, "get:GetBook")
	web.Router("/books/:id", &controllers.SearchController{}, "get:LibraryBook")
	web.Router("/books/:id", &controllers.SearchController{}, "put:UpdateBook")
	web.Router("/books/:id", &controllers.SearchController{}, "delete:DeleteBook")
	web.Router("/books/search/value", &controllers.SearchController{}, "post:Search")
}
