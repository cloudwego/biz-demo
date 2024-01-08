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

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	res, err := c.Echo(context.Background(), &pbapi.Request{Message: "hello"})
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Printf("%v", res)
}
