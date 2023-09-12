package dao

import "wusthelper-mp-gateway/app/model"

const (
	_getStudentBySidSql = "select * from student where sid = ? and deleted = 0"
	_getSidSql          = "select sid from user_basic where oid = ? and deleted = 0"
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
