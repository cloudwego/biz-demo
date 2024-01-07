package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `gorm:"varchar(64),not null"`
}

func (u User) TableName() string {
	return "user"
}
