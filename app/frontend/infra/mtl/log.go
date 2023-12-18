package mtl

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"os"
)

func initLog() {
	// use zap in production mode
	if os.Getenv("GO_ENV") == "online" {
		log := hertzzap.NewLogger()
		hlog.SetLogger(log)
		hlog.SetOutput(os.Stdout)
		hlog.SetLevel(hlog.LevelTrace)
	}
}
