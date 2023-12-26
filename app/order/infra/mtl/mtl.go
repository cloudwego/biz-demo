package mtl

import "sync"

var once sync.Once

func InitMtl() {
	once.Do(
		func() {
			initTracing()
			initMetric()
			initLog()
		},
	)
}
