package mysql

import (
	"fmt"
	"os"

	"github.com/baiyutang/gomall/app/cart/biz/model"
	"github.com/baiyutang/gomall/app/cart/conf"
	"github.com/baiyutang/gomall/app/cart/infra/mtl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		DB.AutoMigrate(
			&model.Cart{},
		)
	}
}
