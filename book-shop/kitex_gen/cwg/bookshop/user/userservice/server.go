// Code generated by Kitex v0.4.3. DO NOT EDIT.
package userservice

import (
	user "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user.UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
