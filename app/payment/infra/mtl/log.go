package mtl

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"os"
)

func initLog() {
	// use zap in production mode
	if os.Getenv("GO_ENV") == "online" {
		log := kitexzap.NewLogger()
		klog.SetLogger(log)
		klog.SetOutput(os.Stdout)
		klog.SetLevel(klog.LevelTrace)
	}
}
