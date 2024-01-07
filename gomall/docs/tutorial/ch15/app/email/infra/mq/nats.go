package mq

import (
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
}
