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

package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("localhost:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		panic(err)
	}

	res, err := cli.Echo(context.Background(), &api.Request{
		Message: "hello",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", res)
}
