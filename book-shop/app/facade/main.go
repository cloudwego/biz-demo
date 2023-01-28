// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"time"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/handlers/handler_item"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/handlers/handler_order"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/handlers/handler_user"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	_ "github.com/cloudwego/biz-demo/book-shop/docs"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/jwt"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func Init() {
	client.Init()

	model.UserAuthMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(conf.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					conf.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar model.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return client.CheckUser(context.Background(), &user.CheckUserReq{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	model.ShopAuthMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(conf.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					conf.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar model.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if loginVar.UserName != conf.ShopLoginName || loginVar.PassWord != conf.ShopLoginPassword {
				return "", jwt.ErrMissingLoginValues
			}

			return conf.ShopLoginName, nil
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

// @title Book-Shop
// @version 1.0
// @description This is a book-shop demo using Hertz and KiteX.

// @contact.name CloudWeGo
// @contact.url https://github.com/cloudwego
// @contact.email conduct@cloudwego.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	Init()
	h := server.Default(server.WithHostPorts(conf.FacadeServiceAddress))
	h.Use(gzip.Gzip(gzip.DefaultCompression))
	pprof.Register(h)

	// user service
	userGroup := h.Group("/user")
	userGroup.POST("/register", handler_user.UserRegister)
	userGroup.POST("/login", handler_user.UserLogin)

	// shop service
	shopGroup := h.Group("/shop")
	shopGroup.POST("/login", handler_user.ShopLogin)

	// item-2b service
	item2BGroup := h.Group("/item2b")
	item2BGroup.Use(model.ShopAuthMiddleware.MiddlewareFunc())
	item2BGroup.POST("/add", handler_item.AddProduct)
	item2BGroup.POST("/edit", handler_item.EditProduct)
	item2BGroup.POST("/del", handler_item.DelProduct)
	item2BGroup.POST("/offline", handler_item.OfflineProduct)
	item2BGroup.POST("/online", handler_item.OnlineProduct)
	item2BGroup.GET("/get", handler_item.GetProduct)
	item2BGroup.POST("/list", handler_item.ListProduct)

	// item-2c service
	item2CGroup := h.Group("/item2c")
	item2CGroup.Use(model.UserAuthMiddleware.MiddlewareFunc())
	item2CGroup.GET("/mget", handler_item.MGetProduct2C)
	item2CGroup.POST("/search", handler_item.SearchProduct)

	// order service
	orderGroup := h.Group("/order")
	orderGroup.Use(model.UserAuthMiddleware.MiddlewareFunc())
	orderGroup.POST("/create", handler_order.CreateOrder)
	orderGroup.POST("/cancel", handler_order.CancelOrder)
	orderGroup.POST("/list", handler_order.ListOrder)
	orderGroup.GET("/get", handler_order.GetOrder)

	url := swagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	h.Spin()
}
