package ecode

func InitEcodeText() {
	texts := map[Code]string{}

	// 通用ecode
	texts[ServerErr] = "服务器未知错误"
	texts[NetworkErr] = "网络异常"
	texts[ParamWrong] = "参数无效"
	texts[FuncNotImpl] = "功能未实现"
	texts[DaoOperationErr] = "Dao操作错误"

	// mp部分
	texts[MpCodeInvalid] = "code无效"
	texts[MpRequestTooFast] = "API调用太频繁"
	texts[MpUserBanned] = "高风险等级用户，小程序登录拦截"
	texts[MpSystemError] = "小程序系统繁忙"

	// rpc部分
	texts[RpcRequestErr] = "Rpc请求错误"

	// 助手主服务响应
	texts[WusthelperAuthErr] = "助手上游：鉴权失败"
	texts[WusthelperTokenInvalid] = "助手上游：token无效"

	// 本科生
	texts[UndergradPasswordWrong] = "本科生：密码错误"
	texts[UndergradUserBanned] = "本科生：用户密码重试过多被封号"
	texts[UndergradPasswordNeedUpdate] = "本科生：密码需要更新（教务处密码已更新）"
	texts[UndergradPasswordNeedModify] = "本科生：密码需要修改"
	texts[UndergradRequestFail] = "本科生：请求失败"

	// 研究生
	texts[GraduatePasswordWrong] = "研究生：密码错误"
	texts[GraduateUserBanned] = "研究生：用户密码重试过多被封号"
	texts[GraduatePasswordNeedUpdate] = "研究生：密码需要更新（教务处密码已更新）"
	texts[GraduatePasswordNeedModify] = "研究生：密码需要修改"
	texts[GraduateRequestFail] = "研究生：请求失败"

	Register(texts)
}
