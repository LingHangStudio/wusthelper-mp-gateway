package ecode

var (
	WusthelperAuthErr      = New(40_000) // 助手上游：鉴权失败
	WusthelperTokenInvalid = New(40_001) // 助手上游：token无效
)
