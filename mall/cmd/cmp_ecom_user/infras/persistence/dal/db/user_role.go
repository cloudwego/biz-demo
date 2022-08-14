package db

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/biz-demo/mall/pkg/conf"
	"gorm.io/gorm"
)

type UserRolePO struct {
	gorm.Model
	UserName string `json:"user_name"`
	Roles    string `json:"roles"`
}

func (userRolePO *UserRolePO) TableName() string {
	return conf.UserRoleTableName
}

func AddUserRole(ctx context.Context, userName string, role user.Role) error {
	records := make([]*UserRolePO, 0)

	insertRecord := &UserRolePO{
		UserName: userName,
	}
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&records).Error
	if err != nil {
		return err
	}
	if len(records) == 0 {
		roleList := []int64{int64(role)}
		rolesBytes, _ := json.Marshal(roleList)
		insertRecord.Roles = string(rolesBytes)
		return DB.WithContext(ctx).Create(insertRecord).Error
	}
	roleListString := records[0].Roles
	roleList := make([]int64, 0)
	_ = json.Unmarshal([]byte((roleListString)), &roleList)
	isExist := false
	for _, roleItem := range roleList {
		if roleItem == int64(role) {
			isExist = true
			break
		}
	}
	if !isExist {
		updateMap := map[string]interface{}{}
		roleList = append(roleList, int64(role))
		rolesBytes, _ := json.Marshal(roleList)
		updateMap["roles"] = string(rolesBytes)
		return DB.WithContext(ctx).Model(&UserRolePO{}).Where("user_name = ?", userName).Updates(updateMap).Error
	}
	return nil
}

func DelUserRole(ctx context.Context, userName string, role user.Role) error {
	records := make([]*UserRolePO, 0)

	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&records).Error
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return nil
	}

	roleListString := records[0].Roles
	roleList := make([]int64, 0)
	updateRoleList := make([]int64, 0)
	_ = json.Unmarshal([]byte((roleListString)), &roleList)
	for _, roleItem := range roleList {
		if roleItem != int64(role) {
			updateRoleList = append(updateRoleList, roleItem)
		}
	}
	rolesBytes, _ := json.Marshal(updateRoleList)
	updateMap := map[string]interface{}{}
	updateMap["roles"] = string(rolesBytes)
	return DB.WithContext(ctx).Model(&UserRolePO{}).Where("user_name = ?", userName).Updates(updateMap).Error
}

func ValidateUserRole(ctx context.Context, userName string, roles []user.Role) (bool, error) {
	record := &UserRolePO{}
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&record).Error
	if err != nil {
		return false, err
	}
	roleListString := record.Roles
	roleList := make([]int64, 0)
	_ = json.Unmarshal([]byte((roleListString)), &roleList)
	roleMap := make(map[int64]bool)
	for _, roleItem := range roleList {
		if roleItem == int64(user.Role_Admin) {
			return true, nil
		}
		roleMap[roleItem] = true
	}
	for _, role := range roles {
		if _, ok := roleMap[int64(role)]; !ok {
			return false, nil
		}
	}
	return true, nil
}
