package service

import (
	"context"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

func (s *Service) GetToken(ctx *context.Context, oid string, platform mp.Platform) (token string, err error) {
	token, err = s.dao.GetToken(ctx, oid)
	if err != nil {
		return "", ecode.DaoOperationErr
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
		return "", ecode.DaoOperationErr
	}

	username, encodedPassword := user.Sid, user.OfficialPwd
	password, err := aes.AesCbcDecryptByHex(encodedPassword, []byte(s.config.Server.PasswordKey), nil)
	if err != nil {
		log.Error("数据库密码解密错误，请检查server配置文件", zap.String("err", err.Error()))
		return "", ecode.ServerErr
	}
	switch user.Type {
	case model.UndergradUser:
		token, _, err = s.UndergradLogin(ctx, username, string(password), oid, false, platform)
		if err != nil {
			return "", err
		}
	case model.GraduateUser:
		token, _, err = s.GraduateLogin(ctx, username, string(password), oid, false, platform)
		if err != nil {
			return "", err
		}
	default:
		return "", ecode.FuncNotImpl
	}

	return token, nil
}
