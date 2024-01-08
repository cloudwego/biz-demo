// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Product{},
			&model.Category{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')")
			DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
		}
	}
}
