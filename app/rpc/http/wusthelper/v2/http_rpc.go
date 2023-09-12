package v2

import (
	"github.com/go-resty/resty/v2"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/library/ecode"
)

type WusthelperHttpRpc struct {
	client *resty.Client
}

func NewRpcClient(c *conf.WusthelperConf) (rpc *WusthelperHttpRpc) {
	httpClient := resty.New()
	httpClient.SetBaseURL(c.Upstream)
	httpClient.SetHeader("User-Agent", "wusthelper-mp-backend/0.0.1")
	httpClient.SetHeader("Platform", "mp")
	httpClient.SetProxy(c.Proxy)
	httpClient.SetTimeout(c.Timeout)

	rpc = &WusthelperHttpRpc{
		client: httpClient,
	}

	return
}

const (
	success                = 10000
	localSuc               = 11000
	limitRequest           = 99999
	err                    = 11111
	baseParamErr           = 10001
	requestDenied          = 10002
	requestCalledErr       = 10003
	clientAbortErr         = 10004
	authErr                = 20001
	authErrTokenMiss       = 21001
	authErrTokenInvalid    = 21002
	authDecodeForStunumSuc = 10000
	authDecodeForStunumErr = 21101
	//身份验证状态码 end
	//jwc模块状态码 start (3)
	jwcPwdNedUpd                       = 30001
	jwcLoginSuc                        = 10000
	jwcLoginErrJwcErr                  = 30104
	jwcLoginErrJwcFinInfoErr           = 30101
	jwcLoginErrJwcModDefPwd            = 30102
	jwcLoginErrJwcErrInfoErr           = 30103
	jwcLoginErrJwcRetryTooManyTimesErr = 30101
	jwcLoginErrJwcUserWasBannedErr     = 30101
	jwcLoginSucJwcErrLocSuc            = 11000
	jwcGetstuinfoSuc                   = 10000
	jwcGetstuinfoSucJwcErrLocSuc       = 11000
	jwcGetstuinfoErrNoSuchStu          = 30201
	jwcGetcurriculumSuc                = 10000
	jwcGetcurriculumSucJwcErrLocSuc    = 11000
	jwcGetcurriculumErrJwcErr          = 30301
	jwcGetjwcannlistSuc                = 10000
	jwcGetjwcanncntSuc                 = 10000
	jwcGetgradeSuc                     = 10000
	jwcGetgradeErrNeedEvaluateTeaching = 30601
	jwcGetgradeErrJwcErrLocSuc         = 11000
	jwcGetcreditSuc                    = 10000
	jwcGetcreditSucJwcErrLocSuc        = 11000
	jwcTrainingSchemeSuc               = 10000
	jwcTrainingSchemeErr               = 30701
	//jwc模块验证码 end
	//lib状态码 start (4)
	libPwdNedUpd                     = 40001
	libLoginExpire                   = 40002
	libLoginSuc                      = 10000
	libLoginErrLibFinInfoErr         = 40101
	libLoginErrWrongAuthCode         = 40103
	libLoginErrTooMuchTime           = 40104
	libLoginErrLibErrInfoErr         = 40102
	libLoginErrLibErrLocSuc          = 11000
	libGetrenthistorySuc             = 10000
	libGetrenthistorySucLibErrLocSuc = 11000
	libGetcurrentrentSuc             = 10000
	libGetcurrentrentSucLibErrLocSuc = 11000
	libGetrentSucLibErrLocSuc        = 11000
	libGetrentSuc                    = 10000
	libGetlibannlistSuc              = 10000
	libGetlibanncntSuc               = 10000
)

func toEcode(code int) ecode.Code {
	return ecode.RpcUnknownErr
}
