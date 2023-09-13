package v1

import (
	"github.com/gin-gonic/gin"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/service"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
)

var (
	serv *service.Service
)

type apiResp[T any] struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
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

//func response(c *gin.Context, resp any) {
//	c.JSON(http.StatusOK, resp)
//}

func response(c *gin.Context, code ecode.Code, data any) {
	responseWithMsg(c, code, code.Message(), data)
}

func responseWithMsg(c *gin.Context, code ecode.Code, msg string, data any) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
