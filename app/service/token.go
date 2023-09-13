package service

import (
	"context"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
)

func (s *Service) GetToken(ctx *context.Context, oid string, platform mp.Platform) (token string, err error) {
	token, err = s.dao.GetToken(ctx, oid)
	if err != nil {
		return "", err
	}

	if token == "" {
		token, err = s.renewWusthelperToken(ctx, oid, platform)
		if err != nil {
			return "", err
		}
	}

	return
}

// renewWusthelperToken 新登录申请一个助手后端token
func (s *Service) renewWusthelperToken(ctx *context.Context, oid string, platform mp.Platform) (token string, err error) {
	user, err := s.GetUserBasic(oid)
	if err != nil {
		return "", err
	}

	username, password := user.Sid, user.OfficialPwd
	switch user.Type {
	case model.UndergradUser:
		token, _, err = s.UndergradLogin(ctx, username, password, oid, platform)
		if err != nil {
			return "", err
		}
	case model.GraduateUser:
		token, _, err = s.GraduateLogin(ctx, username, password, oid, platform)
		if err != nil {
			return "", err
		}
	default:
		return "", ecode.FuncNotImpl
	}

	return token, nil
}
