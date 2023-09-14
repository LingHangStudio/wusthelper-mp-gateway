package ecode

var (
	GraduatePasswordWrong      = New(60_000) // 研究生：密码错误
	GraduateUserBanned         = New(60_001) // 研究生：用户密码重试过多被封号
	GraduatePasswordNeedUpdate = New(60_002) // 研究生：密码需要更新（教务处密码已更新）
	GraduatePasswordNeedModify = New(60_003) // 研究生：密码需要修改
	GraduateRequestFail        = New(60_004) // 研究生：请求失败
)
