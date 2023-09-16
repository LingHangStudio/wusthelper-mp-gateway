package v2

import (
	"go.uber.org/zap"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

func (w *WusthelperHttpRpc) GraduateLogin(username, password string) (token string, err error) {
	query := map[string]string{
		"stuNum": username,
		"jwcPwd": password,
	}

	resp := new(WusthelperResp[string])
	_, err = w.client.R().
		SetQueryParams(query).
		SetResult(resp).
		Post("/yjs/login")
	if err != nil {
		log.Error("助手rpc上游请求出错", zap.String("err", err.Error()))
		return "", ecode.RpcRequestErr
	}

	if resp.Code != success {
		return "", toEcode(resp.Code)
	}

	return resp.Data, nil
}

func (w *WusthelperHttpRpc) GraduateStudentInfo(token string) (studentInfo *GraduateStudentResp, err error) {
	resp := new(WusthelperResp[GraduateStudentResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/yjs/get-student-info")
	if err != nil {
		log.Error("助手rpc上游请求出错", zap.String("err", err.Error()))
		return nil, ecode.RpcRequestErr
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, nil
}

func (w *WusthelperHttpRpc) GraduateCourses(token string) (courses *[]CourseResp, err error) {
	resp := new(WusthelperResp[[]CourseResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/yjs/get-course")
	if err != nil {
		log.Error("助手rpc上游请求出错", zap.String("err", err.Error()))
		return nil, ecode.RpcRequestErr
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, err
}

func (w *WusthelperHttpRpc) GraduateScores(token string) (scores *[]GraduateScoreResp, err error) {
	resp := new(WusthelperResp[[]GraduateScoreResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/yjs/get-grade")
	if err != nil {
		log.Error("助手rpc上游请求出错", zap.String("err", err.Error()))
		return nil, ecode.RpcRequestErr
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, err
}

func (w *WusthelperHttpRpc) GraduateTrainingPlan(token string) (html string, err error) {
	resp := new(WusthelperResp[string])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/yjs/get-scheme")
	if err != nil {
		log.Error("助手rpc上游请求出错", zap.String("err", err.Error()))
		return "", ecode.RpcRequestErr
	}

	if resp.Code != success {
		return "", toEcode(resp.Code)
	}

	return resp.Data, err
}
