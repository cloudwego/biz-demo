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

package conf

const (
	UserTableName    = "t_user"
	ProductTableName = "t_product"
	OrderTableName   = "t_order"

	SecretKey   = "secret key"
	IdentityKey = "id"

	ShopLoginName     = "admin"
	ShopLoginPassword = "123"

	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "127.0.0.1:2379"
	ESAddress       = "http://localhost:9200"
	RedisAddress    = "127.0.0.1:6379"

	RedisConnPoolSize = 20

	RedisKey_User = "user-"

	ProductESIndex = "product"

	UserRpcServiceName   = "cwg.bookshop.user"
	OrderRpcServiceName  = "cwg.bookshop.order"
	ItemRpcServiceName   = "cwg.bookshop.item"
	UserServiceAddress   = "127.0.0.1:8889"
	OrderServiceAddress  = "127.0.0.1:8890"
	ItemServiceAddress   = "127.0.0.1:8891"
	FacadeServiceAddress = "127.0.0.1:8080"
)
