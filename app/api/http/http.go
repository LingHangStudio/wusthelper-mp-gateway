package http

import (
	"github.com/gin-gonic/gin"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/middleware/auth"
	"wusthelper-mp-gateway/app/service"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/ecode/resp"
	"wusthelper-mp-gateway/library/log"
	"wusthelper-mp-gateway/library/token"
)

var (
	serv *service.Service
	jwt  *token.Token
)

func NewEngine(c *conf.Config, baseUrl string) *gin.Engine {
	engine := gin.Default()
	rootRouter := engine.RouterGroup.Group(baseUrl)
	//rootRouter.Use(gin.LoggerWithWriter(*log.DefaultWriter().))

	setupOuterRouter(rootRouter)

	serv = service.New(c)
	jwt = token.New(c.Server.TokenSecret, c.Server.TokenTimeout)

	return engine
}

func setupOuterRouter(group *gin.RouterGroup) {
	mpCommon := group.Group("/")
	{
		mpCommon.POST("/login", mpLogin)
		mpCommon.POST("/decodeToken", auth.UserTokenCheck, mpDecodeToken)
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
		log.Error("获取oid参数失败, oid为空")
		return "", ecode.ParamWrong
	}

	oid, ok := _oid.(string)
	if !ok {
		log.Error("获取oid参数失败, oid转换失败")
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

	c.JSON(200, resp)
}

func toResponseCode(code ecode.Code) (respCode int, msg string) {
	switch code {
	case ecode.UndergradPasswordWrong:
		return resp.UndergradLoginPasswordWrong, "password wrong"
	case ecode.UndergradPasswordNeedUpdate:
		return resp.UndergradLoginPasswordWrong, "password wrong"
	case ecode.UndergradPasswordNeedModify:
		return resp.UndergradLoginPasswordWrong, "password wrong"
	case ecode.GraduatePasswordWrong:
		return resp.GraduateLoginPasswordWrong, "password wrong"
	case ecode.GraduatePasswordNeedUpdate:
		return resp.GraduateLoginPasswordWrong, "password wrong"
	case ecode.GraduatePasswordNeedModify:
		return resp.GraduateLoginPasswordWrong, "password wrong"
	case ecode.WusthelperTokenInvalid:
		return resp.UndergradNeedRelogin, "token invalid"
	}

	return code.Code(), code.Message()
}
