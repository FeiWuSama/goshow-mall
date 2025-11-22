package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/config"
	"workspace-goshow-mall/router"
	"workspace-goshow-mall/utils/logger"
)

func main() {
	conf := config.InitConfig()
	logger.SetLevel(conf.Server.LogLevel)
	logger.Info("starting...", zap.Any("c", conf))
	db, err := initMysql(&conf.MySql)
	if err != nil {
		panic(err)
	}
	logger.Info("mysql connect successfully")
	redisClient, err := initRedis(&conf.Redis)
	if err != nil {
		panic(err)
	}
	logger.Info("redisClient connect successfully")
	startServer(conf, db, redisClient).Run()
}

func startServer(conf config.Config, db *gorm.DB, redisClient *redis.Client) *router.App {
	newAdaptor := adaptor.NewAdaptor(conf, db, redisClient)
	return router.NewApp(conf.Server.Port, router.NewRouter(newAdaptor, conf, func() error {
		err := func() error {
			pingDb, err := db.DB()
			if err != nil {
				panic(err)
			}
			return pingDb.Ping()
		}()
		if err != nil {
			return errors.New("mysql connect failed")
		}
		ctx := context.Background()
		return redisClient.Ping(ctx).Err()
	}))
}

func initMysql(mysqlConfig *config.MySql) (*gorm.DB, error) {
	mysqlConfig.MaxIdleConn = lo.Max([]int{mysqlConfig.MaxIdleConn + 1, 5})
	mysqlConfig.MaxOpenConn = lo.Max([]int{mysqlConfig.MaxOpenConn + 1, 10})
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDb.Ping(); err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	sqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	return db, nil
}

func initRedis(redisConfig *config.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:           fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password:       redisConfig.Password,
		DB:             redisConfig.Db,
		MaxIdleConns:   redisConfig.MaxIdle,
		MaxActiveConns: redisConfig.MaxOpen,
	})
	ctx := context.Background()
	if r, _ := redisClient.Ping(ctx).Result(); r != "PONG" {
		return nil, errors.New("redis connect failed")
	}
	return redisClient, nil
}
