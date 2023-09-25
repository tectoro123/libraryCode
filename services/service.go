package services

import (
	"BookStore/db"
	"BookStore/models"
	"fmt"
)

func DbData(x []map[string]interface{}) {

	fmt.Println(x)

	tabA := []models.Book{}
	for _, eachrecord := range x {

		details := models.Book{}
		details.ID = eachrecord["id"].(float64)
		details.Title = eachrecord["title"].(string)
		details.Publication_year = eachrecord["publication_year"].(float64)
		details.Author = eachrecord["author"].(string)
		details.Price = eachrecord["price"].(float64)
		tabA = append(tabA, details)
	}
	db.Connection.Save(&tabA)
	fmt.Println(tabA)

}
