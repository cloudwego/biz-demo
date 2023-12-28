package mtl

import (
	"sync"

	"github.com/cloudwego/hertz/pkg/route"
)

var once sync.Once

var Hooks []route.CtxCallback

func InitMtl() {
	once.Do(
		func() {
			Hooks = append(Hooks, InitTracing(), initMetric())
			initLog()
		},
	)
}
