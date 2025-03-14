// Code generated by Kitex v0.9.1. DO NOT EDIT.
package submit

import (
	server "github.com/cloudwego/kitex/server"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler submit.Submit, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler submit.Submit, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
