package service

import (
	"time"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/dao"
	v2 "wusthelper-mp-gateway/app/rpc/http/wusthelper/v2"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
)

const (
	wusthelperTokenExpiration = time.Hour * 24
)

type Service struct {
	config *conf.Config

	dao *dao.Dao

	rpc *v2.WusthelperHttpRpc
	mp  *mp.MimiProgram
}

func New(c *conf.Config) (service *Service) {
	service = &Service{
		config: c,
		dao:    dao.New(c),
		rpc:    v2.NewRpcClient(&c.Wusthelper),
		mp:     mp.New(c),
	}

	return
}
