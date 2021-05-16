package main

import (
	"log"
	"net/http"
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	sqlHandler, err := database.NewSqlClient()
	if err != nil {
		log.Fatal(err)
	}
	sqlHandler.Conn.LogMode(true)
	defer sqlHandler.Conn.Close()

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	userHandler := handler.NewUserHandler(*sqlHandler)
	r.POST("/user", func(c *gin.Context) { userHandler.RegisterUser(c) })

	// //CRUDのREST API を作成
	// r.POST("/user", func(c *gin.Context) {
	// 	user := User{}

	// 	err := c.BindJSON(&user)
	// 	if err != nil {
	// 		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	// 	}
	// 	db.NewRecord(user)
	// 	db.Create(&user)
	// 	if db.NewRecord(user) == false {
	// 		c.JSON(http.StatusOK, user)
	// 	}
	// })
	// //READ
	// //全レコード
	// r.GET("/users", func(c *gin.Context) {
	// 	users := []User{}
	// 	db.Find(&users) // 全レコード
	// 	c.JSON(http.StatusOK, users)
	// })
	// //1レコード
	// r.GET("/user/:id", func(c *gin.Context) {
	// 	user := User{}
	// 	id := c.Param("id")

	// 	db.Where("ID = ?", id).First(&user)
	// 	c.JSON(http.StatusOK, user)
	// })

	r.Run(":8080")
}
