package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	proto "utwoo.com/playground/micro-test/proto/greeter"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	//Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	if err := proto.RegisterGreeterHandler(service.Server(), new(Greeter)); err != nil {
		fmt.Println(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
