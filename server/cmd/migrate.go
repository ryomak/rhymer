package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/ryomak/rhymer/server/util"
	"github.com/ryomak/rhymer/server/model"
)

func main() {
	filenPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	var conf struct {
		DBConfig util.DBConfig `toml:"database"`
	}
	if _, err := toml.DecodeFile(*filenPath, &conf); err != nil {
		panic(err)
	}
	db := util.InitDB(conf.DBConfig)
	if err := model.Migrate(db); err != nil {
		panic(err)
	}
}
