package service

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
)

// Code2Session 验证从小程序前端传来的code并换取用户session信息
func (s *Service) Code2Session(code string, platform mp.Platform) (session *mp.SessionInfo, err error) {
	sessionInfo, err := s.mp.Code2Session(platform, code)
	if err != nil {
		return nil, err
	}

	return sessionInfo, nil
}

func (s *Service) GetUserBasic(oid string) (user *model.UserBasic, err error) {
	user, err = s.dao.GetUserBasic(oid)
	if err != nil {
		return nil, err
	}

	return
}

// RegisterUser 登记用户基本信息，如果相应平台的oid已经存在，则直接返回用户基本信息，否则入库保存
func (s *Service) RegisterUser(platform mp.Platform, oid string) (user *model.UserBasic, err error) {
	switch platform {
	case mp.Wechat:
		user, err = s.dao.GetUserBasic(oid)
	case mp.QQ:
		user, err = s.dao.GetUserBasic(oid)
	}

	if user == nil {
		user, err = s.newBasicUser(platform, oid)
		if err != nil {
			return nil, err
		}
	}

	return
}

// newBasicUser 插入一个新的基本用户信息（只有oid和uid的）
func (s *Service) newBasicUser(platform mp.Platform, oid string) (user *model.UserBasic, err error) {
	user = &model.UserBasic{
		Uid:      uint64(idgen.NextId()),
		Oid:      oid,
		Platform: uint8(platform),
	}

	_, err = s.dao.AddUserBasic(user)
	if err != nil {
		return nil, err
	}

	return
}

// SaveUserBasic 保存用户基本信息，用户不存在时插入，存在时则更新
func (s *Service) SaveUserBasic(oid string, userBasic *model.UserBasic, platform mp.Platform) (err error) {
	has, err := s.dao.HasUser(oid)
	if err != nil {
		return err
	}

	if has {
		userBasic.Platform = uint8(platform)
		_, err = s.dao.UpdateUser(oid, userBasic)
	} else {
		uid := idgen.NextId()
		userBasic.Uid = uint64(uid)
		userBasic.Oid = oid
		_, err = s.dao.AddUserBasic(userBasic)
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SaveWxUserProfile(oid string, profile *model.WxUserProfile) (err error) {
	has, err := s.dao.HasUserProfile(mp.Wechat, oid)
	if err != nil {
		return
	}

	if has {
		_, err = s.dao.UpdateWxUserProfile(oid, profile)
	} else {
		profile.Oid = oid
		_, err = s.dao.AddWxUserProfile(profile)
	}
	if err != nil {
		return
	}

	return nil
}

func (s *Service) SaveQQUserProfile(oid string, profile *model.QQUserProfile) (err error) {
	has, err := s.dao.HasUserProfile(mp.QQ, oid)
	if err != nil {
		return
	}

	profile.Oid = oid
	if has {
		_, err = s.dao.UpdateQQUserProfile(oid, profile)
	} else {
		_, err = s.dao.AddQQUserProfile(profile)
	}
	if err != nil {
		return
	}

	return nil
}

func (s *Service) GetWxUserProfile(oid string) (profile *model.WxUserProfile, err error) {
	profile, err = s.dao.GetWxUserProfile(oid)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetQQUserProfile(oid string) (profile *model.QQUserProfile, err error) {
	profile, err = s.dao.GetQQUserProfile(oid)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) CountTotalUser(ctx *context.Context) (count int64, err error) {
	cacheTotal, _ := s.dao.GetTotalUserCountCache(ctx)
	if cacheTotal > 0 {
		return cacheTotal, nil
	}

	cacheTotal, err = s.dao.CountTotalUser()
	if err != nil {
		return 0, err
	}

	_ = s.dao.StoreTotalUserCountCache(ctx, cacheTotal, time.Hour*1)

	return cacheTotal, nil
}

func (s *Service) CheckOidMatchSid(oid, sid string) (result bool, err error) {
	user, err := s.dao.GetUserBasic(oid)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}

	return user.Sid == sid, nil
}
