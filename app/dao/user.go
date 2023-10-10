package dao

import (
	"go.uber.org/zap"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

const (
	_GetUserBasicSql = "select * from user_basic where oid = ? and deleted = 0"
	_HasUserSql      = "select 1 from user_basic where oid = ? and deleted = 0 limit 1"

	_WxUserProfileSql    = "select * from wx_profile where oid = ? and deleted = 0"
	_HasWxUserProfileSql = "select 1 from wx_profile where oid = ? and deleted = 0 limit 1"

	_QQUserProfileSql    = "select * from qq_profile where oid = ? and deleted = 0"
	_HasQQUserProfileSql = "select 1 from qq_profile where oid = ? and deleted = 0 limit 1"
)

func (d *Dao) GetUserBasic(oid string) (user *model.UserBasic, err error) {
	user = new(model.UserBasic)
	has, err := d.db.SQL(_GetUserBasicSql, oid).Get(user)
	if err != nil {
		log.Error("获取用户基本信息出现错误", zap.String("err", err.Error()))
		return nil, err
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) HasUser(oid string) (has bool, err error) {
	has, err = d.db.SQL(_HasUserSql, oid).Exist()
	if err != nil {
		log.Error("查找用户是否存在时出现错误", zap.String("err", err.Error()))
		return false, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) UpdateUser(oid string, user *model.UserBasic, forceUpdate ...string) (count int64, err error) {
	user.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("oid = ?", oid).
		Update(user)
	if err != nil {
		log.Error("更新用户基本信息时出现错误", zap.String("err", err.Error()))
		return 0, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) AddUserBasic(user *model.UserBasic) (count int64, err error) {
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	count, err = d.db.InsertOne(user)
	if err != nil {
		log.Error("插入用户基本信息时出现错误", zap.String("err", err.Error()))
		return 0, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) GetWxUserProfile(oid string) (user *model.WxUserProfile, err error) {
	user = new(model.WxUserProfile)
	has, err := d.db.SQL(_WxUserProfileSql, oid).Get(user)
	if err != nil {
		log.Error("获取微信用户信息出现错误", zap.String("err", err.Error()))
		return nil, ecode.DaoOperationErr
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) GetQQUserProfile(oid string) (user *model.QQUserProfile, err error) {
	user = new(model.QQUserProfile)
	has, err := d.db.SQL(_QQUserProfileSql, oid).Get(user)
	if err != nil {
		log.Error("获取QQ用户信息出现错误", zap.String("err", err.Error()))
		return nil, ecode.DaoOperationErr
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) HasUserProfile(platform mp.Platform, oid string) (has bool, err error) {
	switch platform {
	case mp.Wechat:
		has, err = d.db.SQL(_HasWxUserProfileSql, oid).Exist()
	case mp.QQ:
		has, err = d.db.SQL(_HasQQUserProfileSql, oid).Exist()
	}

	if err != nil {
		log.Error("查找微信/QQ用户信息是否存在时出现错误", zap.String("err", err.Error()), zap.Int("platform", platform))
		return false, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) AddWxUserProfile(user *model.WxUserProfile) (count int64, err error) {
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	count, err = d.db.InsertOne(user)
	if err != nil {
		log.Error("新增微信用户信息时出现错误", zap.String("err", err.Error()))
		return count, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) AddQQUserProfile(user *model.QQUserProfile) (count int64, err error) {
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	count, err = d.db.InsertOne(user)
	if err != nil {
		log.Error("新增QQ用户信息时出现错误", zap.String("err", err.Error()))
		return count, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) UpdateWxUserProfile(oid string, user *model.WxUserProfile, forceUpdate ...string) (count int64, err error) {
	user.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("oid = ?", oid).
		And("deleted = ?", 0).
		Update(user)
	if err != nil {
		log.Error("更新微信用户信息时出现错误", zap.String("err", err.Error()))
		return count, ecode.DaoOperationErr
	}
	return
}

func (d *Dao) UpdateQQUserProfile(oid string, user *model.QQUserProfile, forceUpdate ...string) (count int64, err error) {
	user.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("oid = ?", oid).
		And("deleted = ?", 0).
		Update(user)
	if err != nil {
		log.Error("更新QQ用户信息时出现错误", zap.String("err", err.Error()))
		return count, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) CountTotalUser() (total int64, err error) {
	total, err = d.db.Table("user_basic").Count()
	if err != nil {
		log.Error("查询用户总数时出现错误", zap.String("err", err.Error()))
		return total, ecode.DaoOperationErr
	}

	return
}
