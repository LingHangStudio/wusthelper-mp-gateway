package v2

func (w *WusthelperHttpRpc) UndergradLogin(username, password string) (token string, err error) {
	query := map[string]string{
		"stuNum": username,
		"jwcPwd": password,
	}

	resp := new(WusthelperResp[string])
	_, err = w.client.R().
		SetQueryParams(query).
		SetResult(resp).
		Get("/jwc/login")
	if err != nil {
		return "", err
	}

	if resp.Code != success {
		return "", toEcode(resp.Code)
	}

	return resp.Data, nil
}

func (w *WusthelperHttpRpc) UndergradStudentInfo(token string) (studentInfo *StudentInfoResp, err error) {
	resp := new(WusthelperResp[StudentInfoResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/jwc/get-student-info")
	if err != nil {
		return nil, err
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, nil
}

func (w *WusthelperHttpRpc) UndergradCourses(term, token string) (courses *[]CourseResp, err error) {
	resp := new(WusthelperResp[[]CourseResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetQueryParam("schoolTerm", term).
		SetResult(resp).
		Get("/jwc/get-curriculum")
	if err != nil {
		return nil, err
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, err
}

func (w *WusthelperHttpRpc) UndergradScores(token string) (scores *[]ScoreResp, err error) {
	resp := new(WusthelperResp[[]ScoreResp])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/jwc/get-grade")
	if err != nil {
		return nil, err
	}

	if resp.Code != success {
		return nil, toEcode(resp.Code)
	}

	return &resp.Data, err
}

func (w *WusthelperHttpRpc) UndergradTrainingPlan(token string) (html string, err error) {
	resp := new(WusthelperResp[string])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetResult(resp).
		Get("/jwc/get-scheme")
	if err != nil {
		return "", err
	}

	if resp.Code != success {
		return "", toEcode(resp.Code)
	}

	return resp.Data, err
}
