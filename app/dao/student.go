package dao

import (
	"time"
	"wusthelper-mp-gateway/app/model"
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
		return "", err
	}

	return sid, nil
}

func (d *Dao) GetStudent(sid string) (student *model.Student, err error) {
	student = new(model.Student)
	err = d.db.SQL(_getStudentBySidSql, sid).Find(student)
	if err != nil {
		return nil, err
	}

	return
}

func (d *Dao) HasStudent(sid string) (result bool, err error) {
	result, err = d.db.SQL(_hasStudentSql, sid).Get()
	if err != nil {
		return false, err
	}

	return
}

func (d *Dao) AddStudent(student *model.Student) (count int64, err error) {
	student.CreateTime = time.Now()
	student.UpdateTime = student.CreateTime

	count, err = d.db.InsertOne(student)
	if err != nil {
		return 0, err
	}

	return
}

func (d *Dao) UpdateStudent(sid string, student *model.Student, forceUpdate ...string) (count int64, err error) {
	student.UpdateTime = time.Now()
	count, err = d.db.
		MustCols(forceUpdate...).
		Where("sid = ?", sid).Where("deleted = ?", false).
		Update(student)

	return 0, err
}
