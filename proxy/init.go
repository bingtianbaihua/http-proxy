package proxy

import (
	"os"

	"github.com/bingtianbaihua/goproxy/log"
)

var cnfg Config

func init() {
	err := cnfg.GetConfig("../config/config.json")
	if err != nil {
		log.Error("can not load config file:%v\n", err)
		os.Exit(-1)
	}
}
