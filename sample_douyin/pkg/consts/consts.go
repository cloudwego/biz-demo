// Copyright 2023 CloudWeGo Authors
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

package consts

const (
	UserTableName       = "user"
	VideoTableName      = "video"
	CommentTableName    = "comment"
	MessageTableName    = "message"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	ApiServiceName      = "douyinapi"
	UserServiceName     = "douyinuser"
	VideoServiceName    = "douyinvideo"
	CommentServiceName  = "douyincomment"
	MessageServiceName  = "message"
	RelationTableName   = "relation"
	FavoriteTableName   = "favorite"
	RelationServiceName = "relation"
	FavoriteServiceName = "douyinfavorite"
	MySQLDefaultDSN     = "douyin:douyin@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	VideoServiceAddr    = ":10000"
	RelationServiceAddr = ":12000"
	CommentServiceAddr  = ":13000"
	FavoriteServiceAddr = ":11000"
	MessageServiceAddr  = ":14000"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:10079"
	DefaultLimit        = 10
	//oss相关信息
	Endpoint = "oss-cn-beijing.aliyuncs.com"
	AKID     = "LTAI5tQ4x1ACnZo5brw92kxo"
	AKS      = "SmEavhOQDQ2lBXBaiognBiLuS9N3K9"
	Bucket   = "douyin-video-9567"
	CDNURL   = "http://aliyun.maomint.cn/"
)

// 头像
var AvatarList map[int]string = map[int]string{
	0: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3uldzkb7ij309q09qjsn.jpg",
	1: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3uldztsvxj309q09qdha.jpg",
	2: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule03d3zj309q09qjsm.jpg",
	3: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule0ckvpj309q09qwfh.jpg",
	4: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule0jgguj309q09qmya.jpg",
	5: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule0vqnhj309q09qwg2.jpg",
	6: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule1a2d3j309q09q0tp.jpg",
	7: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule1j42xj309q09qjsx.jpg",
	8: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule1szakj309q09qta0.jpg",
}

// 背景
var BackgroundList map[int]string = map[int]string{
	0: "https://maomint.maomint.cn/douyin/background/125615ape48gysysgxbx0y.jpg",
	1: "https://maomint.maomint.cn/douyin/background/125620l6lecc441lilqej6.jpg",
	2: "https://maomint.maomint.cn/douyin/background/125631yyvjdud5j5tjm9m1.jpg",
	3: "https://maomint.maomint.cn/douyin/background/index.jpg",
}
