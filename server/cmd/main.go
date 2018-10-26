package main

import (
	"github.com/ryomak/rhymer/server"
	"github.com/ryomak/rhymer/server/util"
	"flag"
)

func init() {
	filenPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	if err := util.LoadConfig(*filenPath); err != nil {
		panic(err)
	}
}

func main() {
	server.Run()
}
