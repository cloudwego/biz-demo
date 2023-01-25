// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "CloudWeGo",
            "url": "https://github.com/cloudwego",
            "email": "conduct@cloudwego.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/item2b/add": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "发布商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "发布商品",
                "parameters": [
                    {
                        "description": "发布商品参数",
                        "name": "addProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2b/del": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "删除商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "删除商品",
                "parameters": [
                    {
                        "description": "商品参数",
                        "name": "delProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OperateProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2b/edit": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "编辑商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "编辑商品",
                "parameters": [
                    {
                        "description": "编辑商品参数",
                        "name": "editProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EditProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2b/get": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "商品ID查询商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "商品ID查询商品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商品ID",
                        "name": "product_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2b/list": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "获取商品列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "获取商品列表",
                "parameters": [
                    {
                        "description": "获取商品列表参数",
                        "name": "listProductReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2b/offline": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "下架商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "下架商品",
                "parameters": [
                    {
                        "description": "商品参数",
                        "name": "offlineProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OperateProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2b/online": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "上架商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "上架商品",
                "parameters": [
                    {
                        "description": "商品参数",
                        "name": "onlineProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OperateProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2c/mget": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "商品ID批量查询商品（2C接口）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块-2C"
                ],
                "summary": "商品ID批量查询商品（2C接口）",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品ID 逗号分隔",
                        "name": "product_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/item2c/search": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "搜索商品（2C接口）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块-2C"
                ],
                "summary": "搜索商品（2C接口）",
                "parameters": [
                    {
                        "description": "搜索商品参数",
                        "name": "searchProductReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SearchProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/order/cancel": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "用户取消订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单模块"
                ],
                "summary": "用户取消订单",
                "parameters": [
                    {
                        "description": "取消订单参数",
                        "name": "cancelOrderReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CancelOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/order/create": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "用户下单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单模块"
                ],
                "summary": "用户下单",
                "parameters": [
                    {
                        "description": "提单参数",
                        "name": "createOrderReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/order/get": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "订单ID查询订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单模块"
                ],
                "summary": "订单ID查询订单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "订单ID",
                        "name": "order_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/order/list": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "订单列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单模块"
                ],
                "summary": "订单列表",
                "parameters": [
                    {
                        "description": "订单列表查询参数",
                        "name": "listOrderReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/shop/login": {
            "post": {
                "description": "商家登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家模块"
                ],
                "summary": "商家登录",
                "parameters": [
                    {
                        "description": "店铺账号信息",
                        "name": "userParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "账号信息",
                        "name": "userParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册信息",
                        "name": "userParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddProductRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "isbn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pic": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "spu_name": {
                    "type": "string"
                },
                "spu_price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "model.CancelOrderReq": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                }
            }
        },
        "model.CreateOrderReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "stock_num": {
                    "type": "integer"
                }
            }
        },
        "model.EditProductRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "isbn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pic": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "string"
                },
                "spu_name": {
                    "type": "string"
                },
                "spu_price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "model.ListOrderReq": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ListProductReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "spu_name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "expire": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "model.OperateProductReq": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.SearchProductReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "spu_name": {
                    "type": "string"
                }
            }
        },
        "model.UserParam": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "TokenAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Book-Shop",
	Description:      "This is a book-shop demo using Hertz and KiteX.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}