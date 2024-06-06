// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
