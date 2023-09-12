package service

import (
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/app/dao"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
)

type Service struct {
	config *conf.Config

	dao *dao.Dao

	mp *mp.MimiProgram
}

func New(c *conf.Config) (service *Service) {
	service = &Service{
		config: c,
		dao:    dao.New(c),
		mp:     mp.New(c),
	}

	return
}
