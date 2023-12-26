package mtl

import (
	"net"
	"net/http"

	"github.com/baiyutang/gomall/app/payment/conf"
	"github.com/cloudwego/kitex/pkg/registry"
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

	_ = r.Register(&registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": conf.GetConf().Kitex.Service},
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(conf.GetConf().Kitex.MetricsPort, nil)
}
