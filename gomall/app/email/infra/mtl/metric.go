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
	"net"
	"net/http"

	"github.com/cloudwego/biz-demo/gomall/app/email/conf"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func initMetric() {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	r, _ := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])

	addr, _ := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.MetricsPort)

	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": conf.GetConf().Kitex.Service},
	}

	_ = r.Register(registryInfo)

	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo) //nolint:errcheck
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(conf.GetConf().Kitex.MetricsPort, nil) //nolint:errcheck
}
