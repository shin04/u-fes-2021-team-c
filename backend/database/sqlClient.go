package database

import (
	"u-fes-2021-team-c/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// NewSqlClient initialize a new sql client.
func NewSqlClient(config *config.Config) (*SqlHandler, error) {
	USER := config.DB_USER
	PASS := config.DB_PASS
	HOST := config.DB_HOST
	DBNAME := config.DB_NAME
	dsn := USER + ":" + PASS + "@tcp(" + HOST + ":3306)/" + DBNAME
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		return nil, err
	}

	return &SqlHandler{db}, nil
}
