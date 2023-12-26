package main

import (
	"flag"

	"github.com/fast/internal/config"
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
}
