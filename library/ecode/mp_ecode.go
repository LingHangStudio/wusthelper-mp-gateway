package ecode

var (
	// code2session认证

	MpCodeInvalid    = New(114514) // code无效
	MpRequestTooFast = New(114514) // API调用太频繁
	MpUserBanned     = New(114514) // 高风险等级用户，小程序登录拦截
	MpSystemError    = New(114514) // 小程序系统繁忙

)
