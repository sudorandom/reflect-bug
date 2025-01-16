package main

import (
	"log"
	"net/http"

	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// Import to make sure the global registry is populated
	_ "github.com/sudorandom/reflect-bug/gen/acme/weather/v1"
	_ "github.com/sudorandom/reflect-bug/gen/acme/weather/v2"
)

func main() {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"acme.group.v1.GroupService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
		// Don't forget timeouts!
	}
	log.Fatalf("err: %s", server.ListenAndServe())
}
