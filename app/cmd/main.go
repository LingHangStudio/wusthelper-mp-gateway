package main

import (
	"fmt"
	v1 "wusthelper-mp-gateway/app/api/http/v1"
	"wusthelper-mp-gateway/app/conf"
)

func main() {
	if confErr := conf.Init(); confErr != nil {
		panic(confErr)
	}

	engine := v1.NewEngine(conf.Conf, conf.Conf.Server.BaseUrl)
	addr := fmt.Sprintf("%s:%d", conf.Conf.Server.Address, conf.Conf.Server.Port)
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}
