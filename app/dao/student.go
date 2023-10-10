package dao

import (
	"go.uber.org/zap"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

const (
	_getStudentBySidSql = "select * from student where sid = ? and deleted = 0"
	_getSidSql          = "select sid from user_basic where oid = ? and deleted = 0"

	_hasStudentSql = "select 1 from student where sid = ? and deleted = 0 limit 1"
)

func (d *Dao) GetSid(oid string) (string, error) {
	sid := ""
	_, err := d.db.SQL(_getSidSql, oid).Get(&sid)
	if err != nil {
		log.Error("oid获取sid错误", zap.String("err", err.Error()))
		return "", ecode.DaoOperationErr
	}

	return sid, nil
}

func (d *Dao) GetStudent(sid string) (student *model.Student, err error) {
	student = new(model.Student)
	err = d.db.SQL(_getStudentBySidSql, sid).Find(student)
	if err != nil {
		log.Error("获取student错误", zap.String("err", err.Error()))
		return nil, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) HasStudent(sid string) (result bool, err error) {
	result, err = d.db.SQL(_hasStudentSql, sid).Exist()
	if err != nil {
		log.Error("查询student是否存在时出现错误", zap.String("err", err.Error()))
		return false, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) AddStudent(student *model.Student) (count int64, err error) {
	student.CreateTime = time.Now()
	student.UpdateTime = student.CreateTime

	count, err = d.db.InsertOne(student)
	if err != nil {
		log.Error("添加student错误", zap.String("err", err.Error()))
		return 0, ecode.DaoOperationErr
	}

	return
}

func (d *Dao) UpdateStudent(sid string, student *model.Student, forceUpdate ...string) (count int64, err error) {
	student.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("sid = ?", sid).Where("deleted = ?", false).
		Update(student)
	if err != nil {
		log.Error("更新student错误", zap.String("err", err.Error()))
		return 0, ecode.DaoOperationErr
	}

	return
}
