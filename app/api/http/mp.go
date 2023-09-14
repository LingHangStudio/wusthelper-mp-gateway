package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

func mpLogin(c *gin.Context) {
	platform := getPlatform(c)
	codeData, err := c.GetRawData()
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	code := string(codeData)
	session, err := serv.Code2Session(code, platform)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	oid := session.Openid
	unionid := session.Unionid
	user := &model.UserBasic{
		Oid:      oid,
		Unionid:  session.Unionid,
		Platform: uint8(platform),
	}
	err = serv.SaveUserBasic(oid, user, platform)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	token := jwt.Sign(oid, unionid)

	response(c, ecode.MpLoginOK, "ok", token)
	c.Next()
	return
}

func mpTotalUser(c *gin.Context) {
	ctx := c.Request.Context()
	count, err := serv.CountTotalUser(&ctx)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	data := map[string]any{
		"date":   time.Now().String(),
		"jwcnum": count,
	}
	response(c, ecode.MpCountUserOk, "ok", data)
}

func mpDecodeToken(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	data := map[string]string{
		// 嗯？你问我为什么一定是wechat？那得问老项目了
		"wechat_openid": oid,
	}

	sid, err := serv.GetSid(oid)
	if err != nil {
		return
	}

	if sid == "" {
		response(c, ecode.MpDecodeTokenNoStudent, "sid of this oid doesn't exists", data)
	} else {
		data["stu_num"] = sid
		response(c, ecode.MpDecodeTokenOk, "ok", data)
	}
}

func mpUserProfileUpload(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	platform := getPlatform(c)
	switch platform {
	case mp.Wechat:
		wxUserProfileUpload(c, oid)
	case mp.QQ:
		qqUserProfileUpload(c, oid)
	}
}

func wxUserProfileUpload(c *gin.Context, oid string) {
	req := new(wxUserProfileUploadReq)
	err := c.BindJSON(req)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	profile := &model.WxUserProfile{
		Oid:      oid,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		City:     req.City,
		Province: req.Province,
		Country:  req.Country,
		Avatar:   req.AvatarUrl,
	}

	err = serv.SaveWxUserProfile(oid, profile)
	if err != nil {
		responseEcode(c, ecode.DaoOperationErr, nil)
		return
	}

	response(c, ecode.MpUserProfileUploadOk, "ok", nil)
	return
}

func qqUserProfileUpload(c *gin.Context, oid string) {
	req := new(qqUserProfileUploadReq)
	err := c.BindJSON(req)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	profile := &model.QQUserProfile{
		Oid:      oid,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		City:     req.City,
		Province: req.Province,
		Country:  req.Country,
		Avatar:   req.AvatarUrl,
	}

	err = serv.SaveQQUserProfile(oid, profile)
	if err != nil {
		responseEcode(c, ecode.DaoOperationErr, nil)
		return
	}

	response(c, ecode.MpUserProfileUploadOk, "ok", nil)
	return
}

func mpGetAdminConfigure(c *gin.Context) {
	ctx := c.Request.Context()
	config, err := serv.GetWusthelperAdminConfigure(&ctx)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	response(c, ecode.MpGetAdminConfigureOk, "ok", config)
}

func mpVersionLog(c *gin.Context) {
	versionLog, err := serv.GetVersionLog()
	if err != nil {
		log.Warn("读取版本日志时出现错误", zap.String("err", err.Error()))
	}

	response(c, ecode.MpGetVersionLogOk, "ok", versionLog)
}
