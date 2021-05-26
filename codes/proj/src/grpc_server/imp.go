package main

import (
	"context"
	"ddk/myproto"
)

type HelloServiceImp struct{}

func (p *HelloServiceImp) Hello(
	ctx context.Context, args *myproto.Test,
) (*myproto.Test, error) {
	reply := &myproto.Test{Name: "hello" + args.GetName()}
	return reply, nil
}
