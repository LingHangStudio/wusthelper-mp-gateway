package service

import (
	"context"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"wusthelper-mp-gateway/app/model"
	rpc "wusthelper-mp-gateway/app/rpc/http/wusthelper/v2"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/log"
)

// UndergradLogin 登录并获取学生信息，同时将学生信息入库保存
func (s *Service) UndergradLogin(ctx *context.Context, username, password string, oid string, updateStudentInfo bool, platform mp.Platform) (wusthelperToken string, student *model.Student, err error) {
	wusthelperToken, err = s.rpc.UndergradLogin(username, password)
	if err != nil {
		return "", nil, err
	}

	err = s.dao.StoreWusthelperTokenCache(ctx, wusthelperToken, oid, wusthelperTokenExpiration)
	if err != nil {
		return "", nil, err
	}

	if !updateStudentInfo {
		return wusthelperToken, nil, nil
	}

	encrypted, err := aes.AesCbcEncryptHex([]byte(password), []byte(s.config.Server.PasswordKey), nil)
	if err != nil {
		log.Error("密码加密错误", zap.String("err", err.Error()))
		encrypted = ""
	}
	userBasic := &model.UserBasic{
		Oid:         oid,
		Sid:         username,
		Type:        model.UndergradUser,
		Platform:    uint8(platform),
		OfficialPwd: encrypted,
	}
	err = s.SaveUserBasic(oid, userBasic, platform)
	if err != nil {
		return "", nil, err
	}

	student, err = s.tokenGetUndergradStudentInfo(wusthelperToken)
	if err != nil {
		return "", nil, err
	}

	return wusthelperToken, student, err
}

func (s *Service) UndergradGetStudentInfo(ctx *context.Context, oid string, platform mp.Platform) (student *model.Student, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return nil, err
	}

	return s.tokenGetUndergradStudentInfo(token)
}

func (s *Service) tokenGetUndergradStudentInfo(token string) (student *model.Student, err error) {
	studentInfo, err := s.rpc.UndergradStudentInfo(token)
	if err != nil {
		return nil, err
	}

	student = &model.Student{
		Sid:     studentInfo.StuNum,
		Name:    studentInfo.StuName,
		College: studentInfo.College,
		Major:   studentInfo.Major,
		Clazz:   studentInfo.Classes,
	}
	err = s.SaveStudent(student.Sid, student)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) UndergradGetCourseTable(ctx *context.Context, oid string, term string, platform mp.Platform) (courses *[]rpc.CourseResp, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return nil, err
	}

	courses, err = s.rpc.UndergradCourses(term, token)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) UndergradGetScore(ctx *context.Context, oid string, platform mp.Platform) (scores *[]rpc.ScoreResp, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return nil, err
	}

	scores, err = s.rpc.UndergradScores(token)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) UndergradGetTrainingPlan(ctx *context.Context, oid string, platform mp.Platform) (html string, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return "", err
	}

	html, err = s.rpc.UndergradTrainingPlan(token)
	if err != nil {
		return "", err
	}

	return
}
