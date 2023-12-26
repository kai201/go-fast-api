package main

import (
	"flag"
	"fmt"

	"github.com/fast-api/internal/config"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "c", "config.yml", "configuration file")
	flag.Parse()
	err := config.Init(configFile)
	if err != nil {
		panic("init config error: " + err.Error())
	}

	fmt.Println("test...\n", config.Show())

}
