package data

import (
	"sdk/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSdkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	// 需要多少数据库，就构造多少数据库连接客户端指针
	rdb *redis.Client
}

// 数据库连接：NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		rdb: rdb,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
