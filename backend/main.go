package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "g2538Rk0"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "team_4"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	db := gormConnect()

	defer db.Close()
	db.LogMode(true)
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	//CRUDのREST API を作成
	r.POST("/user", func(c *gin.Context) {
		user := User{}

		err := c.BindJSON(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(user)
		db.Create(&user)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	})
	//READ
	//全レコード
	r.GET("/users", func(c *gin.Context) {
		users := []User{}
		db.Find(&users) // 全レコード
		c.JSON(http.StatusOK, users)
	})
	//1レコード
	r.GET("/user/:id", func(c *gin.Context) {
		user := User{}
		id := c.Param("id")

		db.Where("ID = ?", id).First(&user)
		c.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}

type User struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Password string `json:"password"`
}
