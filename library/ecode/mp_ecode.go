package ecode

var (
	// code2session认证

	MpCodeInvalid    = New(20_001) // code无效
	MpRequestTooFast = New(20_002) // API调用太频繁
	MpUserBanned     = New(20_003) // 高风险等级用户，小程序登录拦截
	MpSystemError    = New(20_004) // 小程序系统繁忙
)
