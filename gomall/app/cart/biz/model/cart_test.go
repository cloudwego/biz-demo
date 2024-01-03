package model

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var (
		err error
		dsn = "root:root@tcp(127.0.0.1:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local"
	)
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return
	}
	m.Run()
}
