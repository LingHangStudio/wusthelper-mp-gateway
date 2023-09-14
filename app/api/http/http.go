package http

import (
	"github.com/gin-gonic/gin"
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
	routerGroup := engine.RouterGroup.Group(baseUrl)
	setupOuterRouter(routerGroup)

	serv = service.New(c)
	jwt = token.New(c.Server.TokenSecret, c.Server.TokenTimeout)

	return engine
}

func setupOuterRouter(group *gin.RouterGroup) {
	mpCommon := group.Group("/")
	{
		mpCommon.GET("/login", mpLogin)
		mpCommon.GET("/getUserNum", mpTotalUser)
		mpCommon.GET("/decodeToken", mpDecodeToken)
		mpCommon.GET("/userInfo", auth.UserTokenCheck, mpUserProfileUpload)
		mpCommon.GET("/getSetting", mpGetAdminConfigure)
		mpCommon.GET("/updateLog", mpVersionLog)
	}

	undergrad := group.Group("/jwc")
	{
		undergrad.POST("/login", auth.UserTokenCheck, undergradLogin)
		undergrad.POST("/getJWCInfo", auth.UserTokenCheck, undergradGetStudentInfo)
		undergrad.POST("/getSchedule", auth.UserTokenCheck, undergradGetCourseTable)
		undergrad.POST("/getGrade", auth.UserTokenCheck, undergradGetScore)
		undergrad.POST("/getprocessMap", auth.UserTokenCheck, undergradGetTrainingPlan)
	}

	graduate := group.Group("/yjs")
	{
		graduate.POST("/login", auth.UserTokenCheck, graduateLogin)
		graduate.POST("/getJWCInfo", auth.UserTokenCheck, graduateGetStudentInfo)
		graduate.POST("/schedule", auth.UserTokenCheck, graduateGetCourseTable)
		graduate.POST("/score", auth.UserTokenCheck, graduateGetScore)
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
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func toResponseCode(code ecode.Code) (respCode int, msg string) {
	//switch code {
	//case ecode.OK:
	//	return respCode, "ok"
	//}

	return code.Code(), code.Message()
}
