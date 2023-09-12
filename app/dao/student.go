package dao

import "wusthelper-mp-gateway/app/model"

const (
	_getStudentBySid = "select * from student where sid = ? and deleted = 0"
	_getWechatSid    = "select sid from user_basic where wx_oid = ? and deleted = 0"
	_getQQSid        = "select sid from user_basic where qq_oid = ? and deleted = 0"
)

func (d *Dao) getSidForWechatUser(oid string) (string, error) {
	sid := ""
	_, err := d.db.SQL(_getWechatSid, oid).Get(&sid)
	if err != nil {
		return "", err
	}

	return sid, nil
}

func (d *Dao) getSidForQQUser(oid string) (string, error) {
	sid := ""
	_, err := d.db.SQL(_getQQSid, oid).Get(&sid)
	if err != nil {
		return "", err
	}

	return sid, nil
}

func (d *Dao) getStudent(sid string) (student *model.Student, err error) {
	student = new(model.Student)
	err = d.db.SQL(_getStudentBySid, sid).Find(student)
	if err != nil {
		return nil, err
	}

	return
}
