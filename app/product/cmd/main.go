package main

import (
	"github.com/baiyutang/gomall/app/product/internal/server"
	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
)

func main() {
	server.Run()
}
