package dal

import "github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/infras/persistence/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
