package main

import (
	"context"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUserServiceImpl(t *testing.T) {
	Init()
	serviceImpl := &UserServiceImpl{}

	rand.Seed(time.Now().UnixNano())
	name := strconv.FormatInt(int64(rand.Intn(1000)), 10)
	pwd := name

	t.Log("===Create Test Begin===")
	createReq := &user.CreateUserReq{
		UserName: name,
		Password: pwd,
	}
	createResp, err := serviceImpl.CreateUser(context.TODO(), createReq)
	t.Logf("resp: %v\n, err: %v", createResp, err)
	assert.Nil(t, err)
	t.Log("===Create Test Pass===")

	t.Log("===Check Test Begin===")
	checkReq := &user.CheckUserReq{
		UserName: name,
		Password: pwd,
	}
	checkResp, err := serviceImpl.CheckUser(context.TODO(), checkReq)
	t.Logf("resp: %v\n, err: %v", checkResp, err)
	assert.Nil(t, err)
	t.Log("===Check Test Pass===")

	t.Log("===GetUser Test Begin===")
	getReq := &user.MGetUserReq{
		Ids: []int64{checkResp.UserId},
	}
	getResp, err := serviceImpl.MGetUser(context.TODO(), getReq)
	t.Logf("resp: %v\n, err: %v", getResp, err)
	assert.Nil(t, err)
	t.Log("===GetUser Test Pass===")
}
