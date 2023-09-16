package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/middleware/auth"
	"wusthelper-mp-gateway/app/service"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/token"
)

var (
	serv *service.Service
	jwt  *token.Token
)

func NewEngine(c *conf.Config, baseUrl string) *gin.Engine {
	engine := gin.Default()
	rootRouter := engine.RouterGroup.Group(baseUrl)
	rootRouter.Use(gin.LoggerWithWriter(log.Default().Writer()))

	setupOuterRouter(rootRouter)

	serv = service.New(c)
	jwt = token.New(c.Server.TokenSecret, c.Server.TokenTimeout)

	return engine
}

func setupOuterRouter(group *gin.RouterGroup) {
	mpCommon := group.Group("/")
	{
		mpCommon.POST("/login", mpLogin)
		mpCommon.POST("/decodeToken", mpDecodeToken)
		//mpCommon.POST("/userInfo", auth.UserTokenCheck, mpUserProfileUpload)
		// 这个接口踏马居然是没有token的
		mpCommon.POST("/userInfo", mpUserProfileUpload)
		mpCommon.GET("/getUserNum", mpTotalUser)
		mpCommon.GET("/getSetting", mpGetAdminConfigure)
		mpCommon.GET("/updateLog", mpVersionLog)
		mpCommon.GET("/info", mpGetUserInfo)
		mpCommon.GET("/getUnionStatus", mpGetUnionStatus)
	}

	undergrad := group.Group("/jwc")
	{
		undergrad.POST("/login", auth.UserTokenCheck, undergradLogin)
		undergrad.GET("/getJWCInfo", auth.UserTokenCheck, undergradGetStudentInfo)
		undergrad.GET("/getSchedule", auth.UserTokenCheck, undergradGetCourseTable)
		undergrad.GET("/getGrade", auth.UserTokenCheck, undergradGetScore)
		undergrad.GET("/getprocessMap", auth.UserTokenCheck, undergradGetTrainingPlan)
	}

	graduate := group.Group("/yjs")
	{
		graduate.POST("/login", auth.UserTokenCheck, graduateLogin)
		graduate.GET("/getJWCInfo", auth.UserTokenCheck, graduateGetStudentInfo)
		graduate.GET("/schedule", auth.UserTokenCheck, graduateGetCourseTable)
		graduate.GET("/score", auth.UserTokenCheck, graduateGetScore)
	}

	library := group.Group("/library")
	library.GET("/")

	physics := group.Group("/wlsy")
	physics.GET("/")
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

func getOid(c *gin.Context) (string, error) {
	_oid, ok := c.Get("oid")
	if !ok {
		return "", ecode.ParamWrong
	}

	oid, ok := _oid.(string)
	if !ok {
		return "", ecode.ParamWrong
	}

	return oid, nil
}

func responseEcode(c *gin.Context, code ecode.Code, data any) {
	respCode, msg := toResponseCode(code)
	response(c, respCode, msg, data)
}

func response(c *gin.Context, code int, msg string, data any) {
	resp := gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	//// 小程序原版没有msg字段，出错的时候data就作为msg
	//if data == nil {
	//	resp["data"] = msg
	//}

	c.JSON(200, resp)
}

func toResponseCode(code ecode.Code) (respCode int, msg string) {
	return code.Code(), code.Message()
}
