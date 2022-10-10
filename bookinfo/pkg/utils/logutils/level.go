// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package logutils

import (
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
)

type Level string

const (
	LevelTrace  Level = "trace"
	LevelDebug  Level = "debug"
	LevelInfo   Level = "info"
	LevelNotice Level = "notice"
	LevelWarn   Level = "warn"
	LevelError  Level = "error"
	LevelFatal  Level = "fatal"
)

// KitexLogLevel return kitex log level
func (level Level) KitexLogLevel() klog.Level {
	l := Level(strings.ToLower(string(level)))
	switch l {
	case LevelTrace:
		return klog.LevelTrace
	case LevelDebug:
		return klog.LevelDebug
	case LevelInfo:
		return klog.LevelInfo
	case LevelNotice:
		return klog.LevelNotice
	case LevelWarn:
		return klog.LevelWarn
	case LevelError:
		return klog.LevelError
	case LevelFatal:
		return klog.LevelFatal
	default:
		return klog.LevelTrace
	}
}
