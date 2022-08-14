package main

import (
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/infras/persistence/dal"
	user "github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/kitex_gen/cmp/ecom/user/userservice"
	"github.com/cloudwego/biz-demo/mall/pkg/conf"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserRpcServiceName}), // server name
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
