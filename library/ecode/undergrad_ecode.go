package ecode

var (
	UndergradPasswordWrong      = New(50_000) // 本科生：密码错误
	UndergradUserBanned         = New(50_001) // 本科生：用户密码重试过多被封号
	UndergradPasswordNeedUpdate = New(50_002) // 本科生：密码需要更新（教务处密码已更新）
	UndergradPasswordNeedModify = New(50_003) // 本科生：密码需要修改
	UndergradRequestFail        = New(50_004) // 本科生：请求失败
)
