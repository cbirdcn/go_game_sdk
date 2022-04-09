package data

import (
	"context"
	"sdk/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	// "sdk/internal/data/ent"
	// "sdk/internal/data/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
// var ProviderSet = wire.NewSet(NewEntClient, NewGormClient, NewData, NewSdkRepo)
var ProviderSet = wire.NewSet(NewGormClient, NewData, NewSdkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	// 需要多少数据库，就构造多少数据库连接客户端指针
	db *gorm.DB // gorm
	// edb *ent.Client // ent不能支持某些DB操作，比如主键命名为ID以外的名字，所以不建议使用
	rdb *redis.Client
}

// func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
// 	client, err := ent.Open(
// 		conf.Database.Driver,
// 		conf.Database.Source,
// 	)
// 	if err != nil {
// 		log.Fatalf("failed opening connection to db: %v", err)
// 	}
// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// 	return client
// }

func NewGormClient(conf *conf.Data, logger log.Logger) *gorm.DB {
	dsn := conf.Database.Source + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	return db
}

// 数据库连接：NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
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
		db: db,
		rdb: rdb,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
