package service

import (
	"go.uber.org/zap"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/library/log"
)

func (s *Service) GetSid(oid string) (string, error) {
	sid, err := s.dao.GetSid(oid)
	if err != nil {
		log.Error("数据库查询出错", zap.String("err", err.Error()))
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
