package dao

import (
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
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
		return nil, err
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) HasUser(oid string) (has bool, err error) {
	has, err = d.db.SQL(_HasUserSql, oid).Get()
	return
}

func (d *Dao) UpdateUser(oid string, user *model.UserBasic, forceUpdate ...string) (count int64, err error) {
	user.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("oid = ?", oid).
		Update(user)

	return
}

func (d *Dao) AddUserBasic(user *model.UserBasic) (count int64, err error) {
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	count, err = d.db.InsertOne(user)
	return
}

func (d *Dao) GetWxUserProfile(oid string) (user *model.WxUserProfile, err error) {
	user = new(model.WxUserProfile)
	has, err := d.db.SQL(_WxUserProfileSql, oid).Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) GetQQUserProfile(oid string) (user *model.QQUserProfile, err error) {
	user = new(model.QQUserProfile)
	has, err := d.db.SQL(_QQUserProfileSql, oid).Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) HasUserProfile(platform mp.Platform, oid string) (has bool, err error) {
	switch platform {
	case mp.Wechat:
		has, err = d.db.SQL(_HasWxUserProfileSql, oid).Get()
	case mp.QQ:
		has, err = d.db.SQL(_HasQQUserProfileSql, oid).Get()
	}

	return
}

func (d *Dao) AddWxUserProfile(user *model.WxUserProfile) (count int64, err error) {
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	count, err = d.db.InsertOne(user)
	return
}

func (d *Dao) AddQQUserProfile(user *model.QQUserProfile) (count int64, err error) {
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	count, err = d.db.InsertOne(user)
	return
}

func (d *Dao) UpdateWxUserProfile(oid string, user *model.WxUserProfile, forceUpdate ...string) (count int64, err error) {
	user.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("oid = ?", oid).
		And("deleted = ?", 0).
		Update(user)

	return
}

func (d *Dao) UpdateQQUserProfile(oid string, user *model.QQUserProfile, forceUpdate ...string) (count int64, err error) {
	user.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("oid = ?", oid).
		And("deleted = ?", 0).
		Update(user)

	return
}

func (d *Dao) CountTotalUser() (total int64, err error) {
	total, err = d.db.Table("user_basic").Count()

	return
}