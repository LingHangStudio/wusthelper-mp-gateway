package dao

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
	"wusthelper-mp-gateway/app/model"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

const (
	_tokenCacheKey = "wusthelper-mp:token:%s"

	_oidSidCacheKey = "wusthelper-mp:oid:%s:sid"
	_sidOidCacheKey = "wusthelper-mp:sid:%s:oid"

	_totalUserCacheKey = "wusthelper-mp:user:total"

	_adminConfigCacheKey = "wusthelper-mp:admin:config"
)

func (d *Dao) StoreWusthelperTokenCache(c *context.Context, token, oid string, ex time.Duration) error {
	key := fmt.Sprintf(_tokenCacheKey, oid)
	err := d.redis.Set(*c, key, token, ex).Err()
	if err != nil {
		log.Error("缓存助手token出现错误", zap.String("err", err.Error()))
		return ecode.DaoOperationErr
	}

	return nil
}

func (d *Dao) StoreOidSidCache(c *context.Context, oid, sid string, ex time.Duration) error {
	oidSidCacheKey := fmt.Sprintf(_oidSidCacheKey, oid)
	sidOidCacheKey := fmt.Sprintf(_sidOidCacheKey, sid)

	err := d.redis.Set(*c, oidSidCacheKey, sid, ex).Err()
	err = d.redis.Set(*c, sidOidCacheKey, oid, ex).Err()
	if err != nil {
		log.Error("缓存Oid-Sid出现错误", zap.String("err", err.Error()))
		return ecode.DaoOperationErr
	}

	return nil
}

func (d *Dao) GetToken(c *context.Context, oid string) (token string, err error) {
	key := fmt.Sprintf(_tokenCacheKey, oid)
	token, err = d.redis.Get(*c, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		log.Error("获取tokenCache出现错误", zap.String("err", err.Error()))
		return "", ecode.DaoOperationErr
	}

	return
}

func (d *Dao) GetSidForOid(c *context.Context, oid string) (sid string, err error) {
	oidSidCacheKey := fmt.Sprintf(_oidSidCacheKey, oid)
	sid, err = d.redis.Get(*c, oidSidCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		log.Error("获取oid-sid出现错误", zap.String("err", err.Error()))
		return "", ecode.DaoOperationErr
	}

	return sid, nil
}

func (d *Dao) GetTotalUserCountCache(ctx *context.Context) (total int64, err error) {
	total, err = d.redis.Get(*ctx, _totalUserCacheKey).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		log.Error("获取用户总数缓存出现错误", zap.String("err", err.Error()))
		return 0, ecode.DaoOperationErr
	}

	return total, nil
}

func (d *Dao) StoreTotalUserCountCache(ctx *context.Context, count int64, ex time.Duration) (err error) {
	err = d.redis.Set(*ctx, _totalUserCacheKey, count, ex).Err()
	if err != nil {
		log.Error("存储用户总数出现错误", zap.String("err", err.Error()))
		return ecode.DaoOperationErr
	}

	return nil
}

func (d *Dao) IncreaseTotalUserCount(ctx *context.Context) (err error) {
	err = d.redis.Incr(*ctx, _totalUserCacheKey).Err()
	if err != nil {
		log.Error("用户总数+1出现错误", zap.String("err", err.Error()))
		return ecode.DaoOperationErr
	}

	return nil
}

func (d *Dao) StoreAdminConfigCache(ctx *context.Context, config *model.AdminConfig, ex time.Duration) (err error) {
	data, err := jsoniter.Marshal(config)
	if err != nil {
		log.Error("序列化管理端配置出现错误", zap.String("err", err.Error()))
		return ecode.DaoOperationErr
	}

	err = d.redis.Set(*ctx, _adminConfigCacheKey, string(data), ex).Err()
	if err != nil {
		log.Error("缓存管理端配置出现错误", zap.String("err", err.Error()))
		return ecode.DaoOperationErr
	}

	return
}

func (d *Dao) GetAdminConfigCache(ctx *context.Context) (config *model.AdminConfig, err error) {
	result, err := d.redis.Get(*ctx, _adminConfigCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		log.Error("获取管理端配置出现错误", zap.String("err", err.Error()))
		return nil, ecode.DaoOperationErr
	}

	config = new(model.AdminConfig)
	err = jsoniter.Unmarshal([]byte(result), config)
	if err != nil {
		log.Error("反序列化管理端配置出现错误", zap.String("err", err.Error()))
		return nil, ecode.DaoOperationErr
	}

	return config, nil
}
