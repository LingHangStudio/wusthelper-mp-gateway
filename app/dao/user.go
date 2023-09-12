package dao

import "wusthelper-mp-gateway/app/model"

const (
	_WxUserSql = "select * from user_basic where wx_oid = ? and deleted = 0 limit 1"
	_QQUserSql = "select * from user_basic where qq_oid = ? and deleted = 0 limit 1"

	_WxUserProfileSql = "select * from wx_profile where oid = ? and deleted = 0 limit 1"
	_QQUserProfileSql = "select * from qq_profile where oid = ? and deleted = 0 limit 1"
)

func (d *Dao) FindWxUserBasic(oid string) (user *model.UserBasic, err error) {
	user = new(model.UserBasic)
	has, err := d.db.SQL(_WxUserSql, oid).Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) FindQQUserBasic(oid string) (user *model.UserBasic, err error) {
	user = new(model.UserBasic)
	has, err := d.db.SQL(_QQUserSql, oid).Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}

	return
}

func (d *Dao) AddWxUserBasic(user *model.UserBasic) (count int64, err error) {
	count, err = d.db.InsertOne(user)
	return
}

func (d *Dao) AddQQUserBasic(user *model.UserBasic) (count int64, err error) {
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
