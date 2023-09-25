package models

type Book struct {
	ID               float64 `gorm:"primaryKey"`
	Title            string  `gorm : "type:varchar(40)"`
	Publication_year float64
	Author           string `gorm : "type:varchar(40)"`
	Price            float64
}
