package service

import (
	"wusthelper-mp-gateway/app/model"
)

func (s *Service) GetSid(oid string) (string, error) {
	sid, err := s.dao.GetSid(oid)
	if err != nil {
		return "", err
	}

	return sid, nil
}

func (s *Service) GetStudent(sid string) (student *model.Student, err error) {
	student, err = s.dao.GetStudent(sid)
	if err != nil {
		return nil, err
	}

	return
}

// SaveStudent 保存学生，如果信息存在则更新，不存在则新增插入
func (s *Service) SaveStudent(sid string, student *model.Student) (err error) {
	hasStudent, err := s.dao.HasStudent(sid)
	if err != nil {
		return err
	}

	if hasStudent {
		_, err = s.dao.UpdateStudent(sid, student)
		if err != nil {
			return err
		}
	} else {
		student.Sid = sid
		_, err = s.dao.AddStudent(student)
		if err != nil {
			return err
		}
	}

	return nil
}
