package mtl

import (
	"context"
	"net"
	"net/http"

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
	config.Address = "10.152.183.21:8500"
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient, consul.WithAdditionInfo(&consul.AdditionInfo{
		Tags: []string{"service:frontend"},
	}))

	ip, err := net.ResolveTCPAddr("tcp", "localhost:8090")
	if err != nil {
		hlog.Error(err)
	}
	registryInfo := &registry.Info{Addr: ip, ServiceName: "prometheus", Weight: 1}
	err = r.Register(registryInfo)

	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(":8090", nil)
	return func(ctx context.Context) {
		r.Deregister(registryInfo)
	}
}
