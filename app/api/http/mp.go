package http

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"net/http"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	respCode "wusthelper-mp-gateway/library/ecode/resp"
	"wusthelper-mp-gateway/library/log"
)

func mpLogin(c *gin.Context) {
	platform := getPlatform(c)
	codeData, err := c.GetRawData()
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	requestJson := map[string]string{}
	err = jsoniter.Unmarshal(codeData, &requestJson)
	code := requestJson["code"]
	if err != nil || code == "" {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	session, err := serv.Code2Session(code, platform)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
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
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	token := jwt.Sign(oid, unionid)

	response(c, respCode.MpLoginOK, "ok", token)
	c.Next()
	return
}

func mpTotalUser(c *gin.Context) {
	ctx := c.Request.Context()
	count, err := serv.CountTotalUser(&ctx)
	if err != nil {
		log.Error("获取用户总数发生错误", zap.String("err", err.Error()))
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	data := map[string]any{
		"date":   time.Now().Format(time.RFC3339),
		"jwcnum": count,
	}
	response(c, respCode.MpCountUserOk, "ok", data)
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
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	if sid == "" {
		response(c, respCode.MpDecodeTokenNoStudent, "sid of this oid doesn't exists", data)
	} else {
		data["stu_num"] = sid
		response(c, respCode.MpDecodeTokenOk, "ok", data)
	}
}

func mpUserProfileUpload(c *gin.Context) {
	//oid, err := getOid(c)
	//if err != nil {
	//	responseEcode(c, ecode.ParamWrong, nil)
	//	return
	//}

	platform := getPlatform(c)
	switch platform {
	case mp.Wechat:
		wxUserProfileUpload(c)
	case mp.QQ:
		qqUserProfileUpload(c)
	}
}

func wxUserProfileUpload(c *gin.Context) {
	req := new(wxUserProfileUploadReq)
	err := c.BindJSON(req)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	oid := req.Oid
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
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	response(c, respCode.MpUserProfileUploadOk, "ok", nil)
	return
}

func qqUserProfileUpload(c *gin.Context) {
	req := new(qqUserProfileUploadReq)
	err := c.BindJSON(req)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	oid := req.Oid
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
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	response(c, respCode.MpUserProfileUploadOk, "ok", nil)
	return
}

func mpGetAdminConfigure(c *gin.Context) {
	ctx := c.Request.Context()
	config, err := serv.GetWusthelperAdminConfigure(&ctx)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	resp := AdminConfigResp{
		Code:        respCode.MpGetAdminConfigureOk,
		TermList:    config.TermList,
		Openadvance: config.Openadvance,
		Schedule:    config.Schedule,
		MenuList:    config.MenuList,
		JumpUnion:   config.JumpUnion,
		Banner:      config.Banner,
		Term:        config.Term,
		ShowNotice:  config.ShowNotice,
		Union:       config.Union,
	}

	c.JSON(http.StatusOK, resp)
}

func mpVersionLog(c *gin.Context) {
	versionLog, err := serv.GetVersionLog()
	if err != nil {
		log.Warn("读取版本日志时出现错误", zap.String("err", err.Error()))
	}

	response(c, respCode.MpGetVersionLogOk, "ok", versionLog)
}

func mpGetUserInfo(c *gin.Context) {
	sid, has := c.GetQuery("stuNum")
	if !has {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	match, err := serv.CheckOidMatchSid(oid, sid)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}
	if !match {
		response(c, respCode.MpGetUserInfoOk, "ok", nil)
		return
	}

	student, err := serv.GetStudent(sid)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}
	if student == nil {
		response(c, respCode.MpGetUserInfoOk, "ok", nil)
		return
	}

	resp := UserInfoResp{
		StuNum:   sid,
		StuName:  student.Name,
		NickName: student.Name,
		College:  student.College,
		Major:    student.Major,
		Classes:  student.Clazz,
		//Birthday:    "",
		//Sex:         "",
		//Nation:      "",
		//NativePlace: "",
		//Phone:       "",
		//Email:       "",
		//QqNum:       "",
		//Wechat:      "",
	}

	response(c, respCode.MpGetUserInfoOk, "ok", resp)
	return
}

func mpGetUnionStatus(c *gin.Context) {
	response(c, respCode.MpGetUnionStatusOk, "ok", 1)
}
