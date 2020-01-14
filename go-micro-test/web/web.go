package main

import (
	"context"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"net/http"
	proto "utwoo.com/playground/micro-test/proto/greeter"
)

var greeterClient proto.GreeterService

func main() {
	registry := consul.NewRegistry()

	service := web.NewService(
		web.Name("web"),
		web.Address(":8088"),
		web.StaticDir("html"),
		web.Registry(registry),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	greeterClient = proto.NewGreeterService("greeter", service.Options().Service.Client())
	service.HandleFunc("/Greeting", Greeting)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	inputString := r.URL.Query().Get("name")

	req := &proto.Request{
		Name: inputString,
	}

	res, err := greeterClient.Hello(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := w.Write([]byte(res.Greeting)); err != nil {
		log.Fatal(err)
	}
}
