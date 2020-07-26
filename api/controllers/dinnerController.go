package controllers

import (
	"app/models/entity"
	"app/models/model"
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	dinners := model.GetAllDinner()

	c.JSON(http.StatusOK, dinners)
}

func Create(c *gin.Context) {
	text := c.PostForm("text")
	category_id := c.PostForm("category_id")
	int_category_id, _ := strconv.Atoi(category_id)
	dinner := entity.Dinner{
		Text:       text,
		CategoryID: int_category_id,
	}

	// バリデーション(github.com/go-ozzo/ozzo-validation)
	err := validation.ValidateStruct(&dinner,
		validation.Field(&dinner.Text, validation.Required),
		validation.Field(&dinner.CategoryID, validation.Required, validation.Max(5)),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		fmt.Println(err)
		return
	}

	model.CreateDinners(&dinner)
}

func Show(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	dinner := model.GetDinnerById(int_id)

	c.JSON(http.StatusOK, dinner)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	model.DeleteDinnerById(int_id)
}
