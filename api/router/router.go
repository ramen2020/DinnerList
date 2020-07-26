package router

import (
	"app/controllers"
	"app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	

)

// Router ルーティング
func Router(dbConn *gorm.DB) {
	r := gin.Default()

	r.POST("/signup", controllers.Signup) // signup
	r.POST("/login", controllers.Login)   // login
	r.GET("/todo", controllers.Index)
	r.Use(middleware.CORS)

	dinner := r.Group("/dinner", middleware.FirebaseAuth)
	{
		dinner.GET("", controllers.Index)              // 一覧
		dinner.POST("/create", controllers.Create)     // 新規作成
		dinner.GET("/:id", controllers.Show)           // 詳細
		dinner.POST("/delete/:id", controllers.Delete) // 削除
	}

	r.Run(":8010")
}

