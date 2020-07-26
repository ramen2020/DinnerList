package model

import (
	"app/models/db"
	"app/models/entity"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAllDinner() []entity. {
	var dinners []entity.Dinner
	db.ConnectDb().Find(&dinners)

	return dinners
}

func CreateDinner(setDinner *entity.Dinner) {
	db.ConnectDb().Create(&setDinner)
}

func GetDinnerById(id int) entity.Dinner {
	var dinner entity.Dinner
	db.ConnectDb().First(&dinner, "id = ?", id)
	fmt.Println(dinner.CreatedAt)

	return dinner
}

func DeleteDinnerById(id int) {
	dinner := []entity.Dinner{}
	db.ConnectDb().Delete(&dinner, id)
}
