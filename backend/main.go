package main

import (
	"log"
	"net/http"
	"os"
	"u-fes-2021-team-c/config"
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	sqlHandler, err := database.NewSqlClient(config)
	if err != nil {
		log.Fatal(err)
	}
	// sqlHandler.Conn.LogMode(true)
	// defer sqlHandler.Conn.Close()

	userRepo := database.NewUserRepository(*sqlHandler)
	studentinfoRepo := database.NewStudentRepository(*sqlHandler)

	r := gin.Default()

	// cors
	r.Use(cors.Default())

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	userHandler := handler.NewUserHandler(userRepo)
	r.POST("/user", func(c *gin.Context) { userHandler.RegisterUser(c) })
	r.GET("/users", func(c *gin.Context) { userHandler.GetAllUsers(c) })
	r.GET("/user", func(c *gin.Context) { userHandler.GetUser(c) })

	studentinfoHandler := handler.NewStudentinfoHandler(studentinfoRepo)
	r.GET("/student_infos", func(c *gin.Context) { studentinfoHandler.GetAllStudentInfo(c) })
	r.GET("/student_info", func(c *gin.Context) { studentinfoHandler.GetStudentInfoByUserId(c) })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
