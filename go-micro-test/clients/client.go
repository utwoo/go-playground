package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	proto "utwoo.com/playground/micro-test/proto/greeter"
)

func main() {
	registry := consul.NewRegistry()

	// Create a new service
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(registry),
	)

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	res, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Greeting)
}
