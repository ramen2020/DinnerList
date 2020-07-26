package entity

import "github.com/jinzhu/gorm"

type Dinner struct {
	gorm.Model
	Text       string `json:"text"		   gorm:"type:varchar(500),not null"`
	CategoryID int    `json:"category_id"	gorm:"int(5),not null"`
}
