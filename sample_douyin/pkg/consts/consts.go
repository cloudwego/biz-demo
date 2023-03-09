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
	Endpoint = "oss-c**************cs.com"
	AKID     = "LTAI****************92kxo"
	AKS      = "SmEa**************LuS9N3K9"
	Bucket   = "douy******************67"
	CDNURL   = "http://*************.cn/"
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
