package v1

import (
	"github.com/gin-gonic/gin"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/service"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
)

var (
	serv *service.Service
)

type apiResp[T any] struct {
	code int    `json:"code,omitempty"`
	msg  string `json:"msg,omitempty"`
	data T      `json:"data,omitempty"`
}

func NewEngine(c *conf.Config, mountUrl string) *gin.Engine {
	engine := gin.Default()
	routerGroup := engine.RouterGroup.Group(mountUrl)
	setupOuterRouter(c, routerGroup)

	serv = service.New(c)

	return engine
}

func setupOuterRouter(c *conf.Config, group *gin.RouterGroup) {
	undergrad := group.Group("/jwc")
	undergrad.POST("/login", undergradLogin)

	admin := group.Group("/admin")
	admin.GET("/list")
}

func getPlatform(c *gin.Context) mp.Platform {
	platformName := c.GetHeader("Platform")
	switch platformName {
	case "wechat":
		return mp.Wechat
	case "qq":
		return mp.QQ
	}

	return mp.Wechat
}
