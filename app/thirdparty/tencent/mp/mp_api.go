package mp

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

type Platform = int

const (
	Wechat = 0
	QQ     = 1
)

var _http = resty.New()

type MimiProgram struct {
	conf *conf.MpConf
}

func New(c *conf.Config) *MimiProgram {
	mp := &MimiProgram{
		conf: &c.MiniProgram,
	}

	return mp
}

func getMpEcode(code int) *ecode.Code {
	switch code {
	case 40029:
		return &ecode.MpCodeInvalid
	case 45011:
		return &ecode.MpRequestTooFast
	case 40226:
		return &ecode.MpUserBanned
	case -1:
		return &ecode.MpSystemError
	default:
		log.Warn("未知errcode", zap.Int("mpEcode", code))
		return nil
	}
}
