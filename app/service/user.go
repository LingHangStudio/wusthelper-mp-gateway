package service

import (
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
)

func (s *Service) Code2Session(platform mp.Platform, code string) (session string, err error) {
	sessionInfo, err := s.mp.Code2Session(platform, code)
	if err != nil {
		return "", err
	}

	return sessionInfo.SessionKey, nil
}

func (s *Service) CheckUser(platform mp.Platform, oid string) (user *model.UserBasic, err error) {
	switch platform {
	case mp.Wechat:
		user, err = s.dao.FindWxUserBasic(oid)
	case mp.QQ:
		user, err = s.dao.FindQQUserBasic(oid)
	}

	if user == nil {

	}

	return
}
