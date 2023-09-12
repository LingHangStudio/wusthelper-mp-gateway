package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yitter/idgenerator-go/idgen"
	v1 "wusthelper-mp-gateway/app/api/http/v1"
	"wusthelper-mp-gateway/app/conf"
)

func main() {
	if confErr := conf.Init(); confErr != nil {
		panic(confErr)
	}

	if conf.Conf.Server.Env == conf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	// 暂时写死workerId为8，要改以后机器扩容了再说
	options := idgen.NewIdGeneratorOptions(8)
	idgen.SetIdGenerator(options)

	engine := v1.NewEngine(conf.Conf, conf.Conf.Server.BaseUrl)
	addr := fmt.Sprintf("%s:%d", conf.Conf.Server.Address, conf.Conf.Server.Port)
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}
