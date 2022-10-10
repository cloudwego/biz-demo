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
