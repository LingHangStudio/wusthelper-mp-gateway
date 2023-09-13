package dao

import (
	"context"
	"fmt"
	"time"
)

const (
	_tokenCacheKey = "wusthelper-mp:token:%s"

	_oidSidCacheKey = "wusthelper-mp:oid:%s:sid"
	_sidOidCacheKey = "wusthelper-mp:sid:%s:oid"
)

func (d *Dao) StoreWusthelperTokenCache(c *context.Context, token, oid string, expiration time.Duration) error {
	key := fmt.Sprintf(_tokenCacheKey, oid)
	err := d.redis.Set(*c, key, token, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) StoreOidSidCache(c *context.Context, oid, sid string, expiration time.Duration) error {
	oidSidCacheKey := fmt.Sprintf(_oidSidCacheKey, oid)
	sidOidCacheKey := fmt.Sprintf(_sidOidCacheKey, sid)

	err := d.redis.Set(*c, oidSidCacheKey, sid, expiration).Err()
	err = d.redis.Set(*c, sidOidCacheKey, oid, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) GetToken(c *context.Context, oid string) (token string, err error) {
	key := fmt.Sprintf(_tokenCacheKey, oid)
	token, err = d.redis.Get(*c, key).Result()
	if err != nil {
		return "", err
	}

	return
}

func (d *Dao) GetSidForOid(c *context.Context, oid string) (sid string, err error) {
	oidSidCacheKey := fmt.Sprintf(_oidSidCacheKey, oid)
	sid, err = d.redis.Get(*c, oidSidCacheKey).Result()
	if err != nil {
		return "", err
	}

	return sid, nil
}
