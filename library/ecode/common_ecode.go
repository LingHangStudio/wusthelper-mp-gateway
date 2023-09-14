package ecode

// All common ecode
var (
	OK = add(0) // 正确

	ServerErr   = add(10_001) // 服务器错误
	NetworkErr  = add(10_002)
	ParamWrong  = add(10_003)
	FuncNotImpl = add(10_004)

	DaoOperationErr = New(10_100)
)
