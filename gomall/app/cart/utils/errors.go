package utils

import "github.com/cloudwego/kitex/pkg/klog"

// MustHandleError log the error info and then exit
func MustHandleError(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}

// ShouldHandleError log the error info
func ShouldHandleError(err error) {
	if err != nil {
		klog.Error(err)
	}
}
