package main

import (
	"flag"
	"time"

	"github.com/himidori/cf-dynamic-dns/cloudflare"
	"github.com/himidori/cf-dynamic-dns/config"
)

var (
	workers           int
	configPath        string
	updateIntervalSec int
)

func init() {
	flag.IntVar(&workers, "workers", 1, "amount of workers for concurrent processing")
	flag.IntVar(&updateIntervalSec, "interval", 30, "update interval")
	flag.StringVar(&configPath, "config", "./config.json", "path to configuration file")
	flag.Parse()
}

func main() {
	config, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}

	for {
		cloudflare.NewCFRequester(config, workers).Run()
		time.Sleep(time.Second * time.Duration(updateIntervalSec))
	}
}
