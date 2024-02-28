// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"log"
	"net"

	"github.com/cloudwego/biz-demo/book-shop/app/item/infras"
	item "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item/itemservice"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	infras.Init()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.ItemServiceAddress)
	if err != nil {
		panic(err)
	}
	svr := item.NewServer(new(ItemServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ItemRpcServiceName}), // server name
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
