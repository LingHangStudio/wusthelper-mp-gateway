package mp

import (
	"go.uber.org/zap"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

const (
	// Code2Session api
	_wechatCode2SessionApi = "https://api.weixin.qq.com/sns/jscode2session"
	_qqCode2SessionApi     = "https://api.q.qq.com/sns/jscode2session"
)

// Code2Session 接口，验证小程序前端传来的code并获取session
func (m *MimiProgram) Code2Session(platform Platform, code string) (*SessionInfo, error) {
	var appid, secret, api string
	if platform == Wechat {
		appid = m.conf.Wechat.AppID
		secret = m.conf.Wechat.Secret
		api = _wechatCode2SessionApi
	} else {
		appid = m.conf.QQ.AppID
		secret = m.conf.QQ.Secret
		api = _qqCode2SessionApi
	}

	query := map[string]string{
		"appid":      appid,
		"secret":     secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	resp := Code2SessionResp{}
	_, err := _http.R().EnableTrace().
		SetQueryParams(query).
		SetResult(&resp).
		Get(api)
	if err != nil {
		log.Error("[Code2Session]请求小程序上游出现异常", zap.String("err", err.Error()))
		return nil, ecode.NetworkErr
	} else if resp.Errcode != 0 {
		log.Debug("[Code2Session]请求小程序上游不成功",
			zap.Int("err", resp.Errcode),
			zap.String("mpErrMsg", resp.Errmsg),
		)

		mpErr := getMpEcode(resp.Errcode)
		return nil, mpErr
	}

	result := &SessionInfo{
		SessionKey: resp.SessionKey,
		Unionid:    resp.Unionid,
		Openid:     resp.Openid,
	}
	return result, nil
}
