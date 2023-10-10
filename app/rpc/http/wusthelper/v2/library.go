package v2

func (w *WusthelperHttpRpc) LibraryLogin(token, password string) (err error) {
	resp := new(WusthelperResp[interface{}])
	_, err = w.client.R().
		SetHeader("Token", token).
		SetQueryParam("libPwd", password).
		SetResult(resp).
		Get("/yjs/login")
	if err != nil {
		return err
	}

	if resp.Code != success {
		return toEcode(resp.Code, "LibraryLogin")
	}

	return nil
}
