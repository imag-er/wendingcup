// Code generated by Kitex v0.9.1. DO NOT EDIT.
package board

import (
	server "github.com/cloudwego/kitex/server"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler board.Board, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler board.Board, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
