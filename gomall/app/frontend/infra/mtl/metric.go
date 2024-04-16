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

package mtl

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	"github.com/cloudwego/biz-demo/gomall/common/utils"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/route"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func initMetric() route.CtxCallback {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	config := consulapi.DefaultConfig()
	config.Address = conf.GetConf().Hertz.RegistryAddr
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient, consul.WithAdditionInfo(&consul.AdditionInfo{
		Tags: []string{"service:frontend"},
	}))

	localIp := utils.MustGetLocalIPv4()
	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", localIp, conf.GetConf().Hertz.MetricsPort))
	if err != nil {
		hlog.Error(err)
	}
	registryInfo := &registry.Info{Addr: ip, ServiceName: "prometheus", Weight: 1}
	err = r.Register(registryInfo)
	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(fmt.Sprintf(":%d", conf.GetConf().Hertz.MetricsPort), nil) //nolint:errcheck
	return func(ctx context.Context) {
		r.Deregister(registryInfo) //nolint:errcheck
	}
}
