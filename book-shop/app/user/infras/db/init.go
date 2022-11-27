package db

import (
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	m := DB.Migrator()
	hasUserTable := m.HasTable(&User{})
	if hasUserTable {
		return
	}
	if !hasUserTable {
		if err = m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
}
