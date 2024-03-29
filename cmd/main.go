package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/fast/internal/accessor"
	"github.com/fast/internal/config"
	"github.com/fast/internal/routers"
	"github.com/fast/pkg/logger"
	"github.com/fast/pkg/stat"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "c", "config.yml", "configuration file")
	flag.Parse()
	initializeConfiguration()
	initializeLogger()
	accessor.InitializeDatabase()
	defer accessor.CloseDatabase()
	initializeApplication()
	initializeRPCServer()
}

func shutdown(sigs chan os.Signal, server *http.Server) {
	sig := <-sigs
	logger.Infof("Signal %s received, shutting down server...", sig)
	ctx := context.Background()

	// Shutdown http server
	err := server.Shutdown(ctx)

	if err != nil {
		logger.Errorf("Failed to shutdown server: %s", err)
	}
	close(sigs)
}

func initializeApplication() {

	router := routers.NewRouter()

	conf := config.Instance
	address := net.JoinHostPort(conf.HTTP.Host, strconv.Itoa(conf.HTTP.Port))
	server := &http.Server{Addr: address, Handler: router}

	sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go shutdown(sigs, server)

	defer func() {
		<-sigs
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Errorf("api run failed %s %s %s", err.Error(), "address", address) //nolint
	}
}

func initializeRPCServer() {
	fmt.Printf("Rpc Test....")
}

func initializeConfiguration() {
	err := config.Init(configFile)
	if err != nil {
		panic("init config error: " + err.Error())
	}

	conf := config.Instance

	if conf.Server.EnableStat {
		stat.Init(
			stat.WithLog(logger.Get()),
			stat.WithAlarm(), // invalid if it is windows, the default threshold for cpu and memory is 0.8, you can modify them
		)
	}
}

func initializeLogger() {
	conf := config.Instance
	// initializing log
	_, err := logger.Init(
		logger.WithLevel(conf.Logger.Level),
		logger.WithFormat(conf.Logger.Format),
		logger.WithSave(conf.Logger.IsSave),
	)

	if err != nil {
		panic(err)
	}
}
