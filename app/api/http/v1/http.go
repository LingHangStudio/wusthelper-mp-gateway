package v1

import (
	"github.com/gin-gonic/gin"
	"wusthelper-mp-gateway/app/conf"
)

func NewEngine(c *conf.Config, mountUrl string) *gin.Engine {
	engine := gin.Default()
	routerGroup := engine.RouterGroup.Group(mountUrl)
	setupRouter(c, routerGroup)

	return engine
}

func setupRouter(c *conf.Config, group *gin.RouterGroup) {
	public := group.Group("/recruit")
	public.POST("/submit")

	admin := group.Group("/admin")
	admin.GET("/list")
}
