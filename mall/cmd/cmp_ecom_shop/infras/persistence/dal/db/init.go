package db

import (
	"github.com/cloudwego/biz-demo/mall/pkg/conf"
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
	if m.HasTable(&ShopPO{}) {
		return
	}
	if err = m.CreateTable(&ShopPO{}); err != nil {
		panic(err)
	}
}
