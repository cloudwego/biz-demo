package db

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/kitex_gen/cmp/ecom/user"
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
	hasUserTable := m.HasTable(&UserPO{})
	hasUserRoleTable := m.HasTable(&UserRolePO{})
	if hasUserTable && hasUserRoleTable {
		return
	}
	if !hasUserTable {
		if err = m.CreateTable(&UserPO{}); err != nil {
			panic(err)
		}
	}
	if !hasUserRoleTable {
		if err = m.CreateTable(&UserRolePO{}); err != nil {
			panic(err)
		}
		// insert admin
		AdminUser := &UserPO{
			UserName: "admin",
			Password: "admin",
		}
		ctx := context.TODO()
		userList, err := QueryUser(ctx, "admin")
		if err != nil {
			panic(err)
		}
		if len(userList) == 0 {
			if err = CreateUser(ctx, []*UserPO{AdminUser}); err != nil {
				panic(err)
			}
		}
		if err = AddUserRole(ctx, "admin", user.Role_Admin); err != nil {
			panic(err)
		}
	}
}
