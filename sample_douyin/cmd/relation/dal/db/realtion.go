package db

import (
	"context"
	"mydouyin/pkg/consts"
	"time"

	"gorm.io/gorm"
)

type Relation struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	FollowId   int64 `json:"follow_id"`
	FollowerId int64 `json:"follower_id"`
}

func (u *Relation) TableName() string {
	return consts.RelationTableName
}

func CreateRelation(ctx context.Context, realtion *Relation) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(realtion).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", realtion.FollowId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", realtion.FollowerId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

func DeleteRelation(ctx context.Context, follow_id, follower_id int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("follow_id = ? AND follower_id = ?", follow_id, follower_id).Delete(&Relation{}).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", follow_id).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", follower_id).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func ValidRelationIfExist(ctx context.Context, follow_id, follower_id int64) (bool, error) {
	res := make([]*Relation, 0)
	if err := DB.WithContext(ctx).Model(&Relation{}).Where("follow_id = ? AND follower_id = ?", follow_id, follower_id).Find(&res).Error; err != nil {
		return false, err
	}
	return len(res) > 0, nil
}

// 通过FollowId查询所有FollowerId
func GetFollowersByFollow(ctx context.Context, follow_id int64) ([]*Relation, error) {
	res := make([]*Relation, 0)
	if err := DB.WithContext(ctx).Where("follow_id = ?", follow_id).Select("follower_id").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 通过FollowerId查询所有FollowId
func GetFollowsByFollower(ctx context.Context, follower_id int64) ([]*Relation, error) {
	res := make([]*Relation, 0)
	if err := DB.WithContext(ctx).Where("follower_id = ?", follower_id).Select("follow_id").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetFriend(ctx context.Context, me int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := DB.WithContext(ctx).Table("relation as a").Distinct("a.follow_id as friend_id").
		Joins("inner join relation as  b on a.follower_id = ? and b.follow_id = ?", me, me).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
