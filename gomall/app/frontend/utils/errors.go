package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MustHandleError(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}
