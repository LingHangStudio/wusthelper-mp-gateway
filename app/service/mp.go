package service

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"os"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/library/log"
)

func (s *Service) GetWusthelperAdminConfigure(ctx *context.Context) (config *model.AdminConfig, err error) {
	config, err = s.dao.GetAdminConfigCache(ctx)
	if err != nil {
		log.Warn("获取管理端配置缓存时出现错误", zap.String("err", err.Error()))
	}

	if config != nil {
		return config, nil
	}

	config, err = s.rpc.GetAdminConfigure()
	if err != nil {
		log.Error("获取管理端配置缓存时出现错误", zap.String("err", err.Error()))
		return nil, err
	}

	err = s.dao.StoreAdminConfigCache(ctx, config, time.Hour*24*3)
	if err != nil {
		log.Warn("保存管理端配置缓存时出现错误", zap.String("err", err.Error()))
	}

	return
}

func (s *Service) GetVersionLog() (versionLogs *[]model.VersionLog, err error) {
	fileLocation := s.config.Server.VersionLogFile
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		return nil, err
	}

	versionLogs = new([]model.VersionLog)
	err = jsoniter.Unmarshal(file, versionLogs)
	if err != nil {
		return nil, err
	}

	return versionLogs, nil
}
