package service

import (
	"time"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/dao"
	v2 "wusthelper-mp-gateway/app/rpc/http/wusthelper/v2"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/token"
)

const (
	wusthelperTokenExpiration = time.Hour * 24
)

type Service struct {
	config *conf.Config

	whToken *token.Token
	jwt     *token.Token

	dao *dao.Dao

	rpc *v2.WusthelperHttpRpc
	mp  *mp.MimiProgram
}

func New(c *conf.Config) (service *Service) {
	service = &Service{
		config:  c,
		dao:     dao.New(c),
		whToken: token.New(c.Wusthelper.TokenKey, c.Wusthelper.Timeout),
		jwt:     token.New(c.Server.TokenSecret, c.Server.TokenTimeout),
		rpc:     v2.NewRpcClient(&c.Wusthelper),
		mp:      mp.New(c),
	}

	return
}
