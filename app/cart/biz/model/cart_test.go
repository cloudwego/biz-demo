package model

import (
	"context"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var (
		err error
		dsn = "root:root@tcp(127.0.0.1:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local"
	)
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return
	}
	m.Run()
}

func TestAdd(t *testing.T) {
	err := Add(db, context.TODO(), &CartItem{UserId: 11, ProductId: 22, Quantity: 3})
	if err != nil {
		t.Error(err)
	}
}

func TestEmpty(t *testing.T) {
	err := Empty(db, context.TODO(), 11)
	if err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {
	_ = Empty(db, context.TODO(), 11)

	err := Add(db, context.TODO(), &CartItem{UserId: 11, ProductId: 22, Quantity: 3})
	if err != nil {
		t.Error(err)
	}

	list, err := GetCartList(db, context.TODO(), 11)
	if err != nil {
		t.Error(err)
	}
	if len(list) == 1 {
		t.Log("ok")
	} else {
		t.Error("not ok")
	}
}
