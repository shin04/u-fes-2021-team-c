package database

import (
	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// NewSqlClient initialize a new sql client.
func NewSqlClient() (*SqlHandler, error) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "team_4"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil, err
	}

	return &SqlHandler{db}, nil
}
