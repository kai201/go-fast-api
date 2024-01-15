package cmd

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/server"
)

type CmdBoot struct {
}

func New() *CmdBoot {
	app := &CmdBoot{}

	var configPath string
	var runtimePath string
	flag.StringVar(&configPath, "configs", "configs", "config yml files path")
	flag.StringVar(&runtimePath, "runtime", "runtime", "runtime log files path")
	flag.Parse()

	app.bootstrap(configPath, runtimePath)

	return app
}

func (g *CmdBoot) bootstrap(configPath, runtimePath string) {

}

func (g *CmdBoot) Listen(routesFn func(r *gin.Engine), rpc interface{}, authFn func() gin.HandlerFunc) {

	address := "0.0.0.0"
	go runHttpServer(address, routesFn, authFn)
	go runRpcServer(address, "tcp", "", rpc, "")
}

func runHttpServer(listenAddress string, routes func(r *gin.Engine), auth func() gin.HandlerFunc) error {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	return engine.Run(listenAddress)
}

func runRpcServer(name, network, address string, obj interface{}, metadata string) error {
	s := server.NewServer()

	if err := s.RegisterName(name, obj, metadata); err != nil {
		return err
	}
	if err := s.Serve(network, address); err != nil {
		return err
	}
	return nil
}
