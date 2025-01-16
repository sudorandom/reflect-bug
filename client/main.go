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
		fmt.Println("Fetching acme.weather.v1.WeatherService (this service imports a file in the same package)")
		files, err := stream.FileContainingSymbol("acme.weather.v1.WeatherService")
		if err != nil {
			log.Fatalf("err: %s", err)
			return
		}
		for _, file := range files {
			fmt.Println(file.GetName(), file.GetPackage())
		}
		// acme/weather/v1/service.proto acme.weather.v1

		// This shows that querying for acme.weather.v1.WeatherService only returns acme/weather/v1/service.proto.
		// It is MISSING acme/weather/v1/types.proto, which is needed to have a full view of how to use this service.
	}

	fmt.Println()

	{
		fmt.Println("Fetching acme.weather.v2.WeatherService (split into multiple packages)")
		files, err := stream.FileContainingSymbol("acme.weather.v2.WeatherService")
		if err != nil {
			log.Fatalf("err: %s", err)
			return
		}
		for _, file := range files {
			fmt.Println(file.GetName(), file.GetPackage())
		}
		// acme/weather/v2/service.proto acme.weather.v2
		// acme/weather/shared/types.proto acme.weather.shared

		// This shows that if you move types.proto to a different package, receive the types.proto FileDescriptor in the response.
	}
}
