package mtl

import (
	"context"
	"fmt"
	"github.com/baiyutang/gomall/app/frontend/conf"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
	config.Address = conf.GetConf().Hertz.RegistryAddr
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient, consul.WithAdditionInfo(&consul.AdditionInfo{
		Tags: []string{"service:frontend"},
	}))

	localIp := utils.LocalIP()
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
	go http.ListenAndServe(fmt.Sprintf(":%d", conf.GetConf().Hertz.MetricsPort), nil)
	return func(ctx context.Context) {
		r.Deregister(registryInfo)
	}
}
