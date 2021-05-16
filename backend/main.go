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
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	userHandler := handler.NewUserHandler(*sqlHandler)
	r.POST("/user", func(c *gin.Context) { userHandler.RegisterUser(c) })
	r.GET("/users", func(c *gin.Context) { userHandler.GetAllUsers(c) })
	r.GET("/user/:id", func(c *gin.Context) { userHandler.GetUser(c) })

	r.Run(":8080")
}
