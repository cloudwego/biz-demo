module github.com/cloudwego/biz-demo/book-shop

go 1.16

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/kitex v0.4.3
	github.com/hertz-contrib/swagger v0.0.0-20220711030440-b6402d4709f0
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
)

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cloudwego/hertz v0.3.2
	github.com/hertz-contrib/jwt v1.0.1
	github.com/jinzhu/copier v0.3.5
	github.com/kitex-contrib/registry-etcd v0.0.0-20221223084757-0d49e7162359
	github.com/olivere/elastic/v7 v7.0.32
	github.com/r3labs/diff/v2 v2.15.1
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)

replace go.etcd.io/etcd/server/v3 => go.etcd.io/etcd/server/v3 v3.5.7
