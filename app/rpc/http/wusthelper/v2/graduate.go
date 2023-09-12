package v2

func (w *WusthelperHttpRpc) GraduateLogin(username, password string) (token string, err error) {
	query := map[string]string{
		"stuNum": username,
		"jwcPwd": password,
	}

	resp := new(WusthelperResp[string])
	_, err = w.client.R().
		SetQueryParams(query).
		SetResult(resp).
		Get("/yjs/login")
	if err != nil {
		return "", err
	}

	if resp.Code != success {
		return "", toEcode(resp.Code)
	}

	return resp.Data, nil
}

func (w *WusthelperHttpRpc) GraduateStudentInfo(token string) (studentInfo *StudentInfoResp, err error) {
	resp := new(WusthelperResp[StudentInfoResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/yjs/get-student-info")
	if err != nil {
		return nil, err
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, nil
}

func (w *WusthelperHttpRpc) GraduateCourses(term, token string) (courses *[]CourseResp, err error) {
	resp := new(WusthelperResp[[]CourseResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetQueryParam("schoolTerm", term).
		SetResult(resp).
		Get("/yjs/get-curriculum")
	if err != nil {
		return nil, err
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
		return nil, err
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
		return "", err
	}

	if resp.Code != success {
		return "", toEcode(resp.Code)
	}

	return resp.Data, err
}
