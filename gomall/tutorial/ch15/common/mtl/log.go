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
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog() {
	var opts []kitexzap.Option
	var output zapcore.WriteSyncer
	if os.Getenv("GO_ENV") != "online" {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = os.Stdout
	} else {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(os.Stdout),
			FlushInterval: time.Minute,
		}
		server.RegisterShutdownHook(func() {
			output.Sync() //nolint:errcheck
		})
	}
	log := kitexzap.NewLogger(opts...)
	klog.SetLogger(log)
	klog.SetLevel(klog.LevelTrace)
	klog.SetOutput(output)
}
