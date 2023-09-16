package service

import (
	"context"
	"github.com/wumansgy/goEncrypt/aes"
	"go.uber.org/zap"
	"wusthelper-mp-gateway/app/model"
	rpc "wusthelper-mp-gateway/app/rpc/http/wusthelper/v2"
	"wusthelper-mp-gateway/app/thirdparty/tencent/mp"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

// GraduateLogin 登录并获取学生信息，同时将学生信息入库保存
func (s *Service) GraduateLogin(ctx *context.Context, username, password string, oid string, updateStudentInfo bool, platform mp.Platform) (wusthelperToken string, student *model.Student, err error) {
	wusthelperToken, err = s.rpc.GraduateLogin(username, password)
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
		Type:        model.GraduateUser,
		Platform:    uint8(platform),
		OfficialPwd: encrypted,
	}
	err = s.SaveUserBasic(oid, userBasic, platform)
	if err != nil {
		return "", nil, err
	}

	student, err = s.tokenGetGraduateStudentInfo(wusthelperToken)
	if err != nil {
		return "", nil, err
	}

	return wusthelperToken, student, err
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
		return nil, ecode.RpcRequestErr
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
		return nil, ecode.RpcRequestErr
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
		return "", ecode.RpcRequestErr
	}

	return
}
