package v2

import (
	"go.uber.org/zap"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

func (w *WusthelperHttpRpc) GetAdminConfigure() (config *model.AdminConfig, err error) {
	resp := new(WusthelperResp[model.AdminConfig])
	_, err = w.adminClient.R().
		SetResult(resp).
		Get("/wusthelper/config")
	if err != nil {
		log.Error("助手rpc上游请求出错", zap.String("err", err.Error()))
		return nil, ecode.RpcRequestErr
	}

	if resp.Code != adminSuccess {
		return nil, toEcode(resp.Code, "GetAdminConfigure")
	}

	return &resp.Data, nil
}
