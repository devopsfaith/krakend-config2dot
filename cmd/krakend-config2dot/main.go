package main

import (
	"flag"
	"log"
	"os"

	dot "github.com/devopsfaith/krakend-config2dot"
	"github.com/luraproject/lura/config"
)

func main() {
	configFile := flag.String("c", "config.json", "path to the krakend config file")
	flag.Parse()

	cfg, err := config.NewParser().Parse(*configFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	dotCfg := dot.ServiceConfig(cfg)
	dotCfg.WriteTo(os.Stdout)
}
