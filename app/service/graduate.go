package service

import (
	"context"
	"wusthelper-mp-gateway/app/model"
	rpc "wusthelper-mp-gateway/app/rpc/http/wusthelper/v2"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
)

// GraduateLogin 登录并获取学生信息，同时将学生信息入库保存
func (s *Service) GraduateLogin(ctx *context.Context, username, password string, oid string, platform mp.Platform) (token string, student *model.Student, err error) {
	token, err = s.rpc.GraduateLogin(username, password)
	if err != nil {
		return "", nil, err
	}

	err = s.dao.StoreWusthelperTokenCache(ctx, token, oid, wusthelperTokenExpiration)
	if err != nil {
		return "", nil, err
	}

	userBasic := &model.UserBasic{
		Oid:         oid,
		Sid:         username,
		Platform:    uint8(platform),
		OfficialPwd: password,
	}
	err = s.SaveUserBasic(oid, userBasic, platform)
	if err != nil {
		return "", nil, err
	}

	student, err = s.tokenGetGraduateStudentInfo(token)
	if err != nil {
		return "", nil, err
	}

	return token, student, err
}

func (s *Service) GraduateGetStudentInfo(ctx *context.Context, oid string, platform mp.Platform) (student *model.Student, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return nil, ecode.DaoOperationErr
	}

	return s.tokenGetGraduateStudentInfo(token)
}

func (s *Service) tokenGetGraduateStudentInfo(token string) (student *model.Student, err error) {
	studentInfo, err := s.rpc.GraduateStudentInfo(token)
	if err != nil {
		return nil, ecode.DaoOperationErr
	}

	student = &model.Student{
		Sid:     studentInfo.StudentNum,
		Name:    studentInfo.Name,
		College: studentInfo.Academy,
		Major:   studentInfo.Specialty,
		Clazz:   studentInfo.Specialty,
	}
	err = s.SaveStudent(student.Sid, student)
	if err != nil {
		return nil, ecode.DaoOperationErr
	}

	return
}

func (s *Service) GraduateGetCourseTable(ctx *context.Context, oid string, platform mp.Platform) (courses *[]rpc.CourseResp, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return nil, err
	}

	courses, err = s.rpc.GraduateCourses(token)
	if err != nil {
		return nil, ecode.RpcUnknownErr
	}

	return
}

func (s *Service) GraduateGetScore(ctx *context.Context, oid string, platform mp.Platform) (scores *[]rpc.GraduateScoreResp, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return nil, err
	}

	scores, err = s.rpc.GraduateScores(token)
	if err != nil {
		return nil, ecode.RpcUnknownErr
	}

	return
}

func (s *Service) GraduateGetTrainingPlan(ctx *context.Context, oid string, platform mp.Platform) (html string, err error) {
	token, err := s.GetToken(ctx, oid, platform)
	if err != nil {
		return "", err
	}

	html, err = s.rpc.GraduateTrainingPlan(token)
	if err != nil {
		return "", ecode.RpcUnknownErr
	}

	return
}
