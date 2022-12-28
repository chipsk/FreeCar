// Code generated by hertz generator.

package main

import (
	"fmt"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/global"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/initialize"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/initialize/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
)

func main() {
	// initialize
	initialize.InitLogger()
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	tracer, cfg := hertztracing.NewServerTracer()
	rpc.Init()
	// create a new server
	h := server.New(
		tracer,
		server.WithHostPorts(fmt.Sprintf(":%d", global.ServerConfig.Port)),
		server.WithRegistry(r, info),
		server.WithHandleMethodNotAllowed(true),
	)
	// use pprof & tracer mw
	pprof.Register(h)
	h.Use(hertztracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
