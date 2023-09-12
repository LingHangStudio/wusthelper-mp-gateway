package dao

import (
	goredis "github.com/redis/go-redis/v9"
	"wusthelper-mp-gateway/app/conf"
	"wusthelper-mp-gateway/library/cache/redis"
	"wusthelper-mp-gateway/library/database"
	"wusthelper-mp-gateway/library/log"
	"xorm.io/xorm"
)

type Dao struct {
	db    *xorm.Engine
	redis *goredis.Client
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db:    database.NewMysql(&c.Database),
		redis: redis.NewRedisClient(&c.Redis),
	}

	return
}

func (d *Dao) Close() {
	dbErr := d.db.Close()
	if dbErr != nil {
		log.Warn("[dao]关闭数据库连接出错")
	}

	redisErr := d.redis.Close()
	if redisErr != nil {
		log.Warn("[dao]关闭redis连接出错")
	}
}
