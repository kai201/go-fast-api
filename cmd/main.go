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
	"time"

	"github.com/fast/internal/config"
	"github.com/fast/internal/routers"
	"github.com/fast/pkg/logger"
	"github.com/fast/pkg/stat"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "c", "config.yml", "configuration file")
	flag.Parse()
	err := config.Init(configFile)
	if err != nil {
		panic("init config error: " + err.Error())
	}

	// initializing log
	_, err = logger.Init(
		logger.WithLevel(config.Instance.Logger.Level),
		logger.WithFormat(config.Instance.Logger.Format),
		logger.WithSave(config.Instance.Logger.IsSave),
	)

	if err != nil {
		panic(err)
	}

	if config.Instance.Server.EnableStat {
		stat.Init(
			stat.WithLog(logger.Get()),
			stat.WithAlarm(), // invalid if it is windows, the default threshold for cpu and memory is 0.8, you can modify them
		)
	}
	// log := logger.Get()
	router := routers.NewRouter()
	address := net.JoinHostPort("0.0.0.0", strconv.Itoa(8080))
	server := http.Server{Addr: address, Handler: router}
	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			// log.Error("api run failed", err, "address", address)
			logger.Errorf("api run failed", err, "address", address)
			os.Exit(1)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// graceful shutdown operation.
	if err := server.Shutdown(ctx); err != nil {
		// log.ZError(context.Background(), "failed to api-server shutdown", err)
		fmt.Println(err.Error())
	}
}
