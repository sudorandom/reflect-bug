package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"

	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
)

func newInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				// If you're also using this client for non-h2c traffic, you may want
				// to delegate to tls.Dial if the network isn't TCP or the addr isn't
				// in an allowlist.
				return net.Dial(network, addr)
			},
			// Don't forget timeouts!
		},
	}
}

func main() {
	client := grpcreflect.NewClient(newInsecureClient(), "http://127.0.0.1:8080")
	stream := client.NewStream(context.Background())
	{
		files, err := stream.FileContainingSymbol("acme.weather.v1.WeatherService")
		if err != nil {
			log.Fatalf("err: %s", err)
			return
		}
		for _, file := range files {
			fmt.Println(file.GetName(), file.GetPackage())
		}
	}
	{
		files, err := stream.FileContainingSymbol("acme.weather.v2.WeatherService")
		if err != nil {
			log.Fatalf("err: %s", err)
			return
		}
		for _, file := range files {
			fmt.Println(file.GetName(), file.GetPackage())
		}
	}
}
