package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本,如果爆红,把包删除即可,重新import
	"github.com/spf13/viper"
	"sync"
)

type RedisHelper struct {
	*redis.Client
}

// 声明一个全局变量
var (
	redisHelper *RedisHelper

	redisOnce sync.Once
	ctx       = context.Background()
)

// RedisInit Init 初始化连接
func RedisInit() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", RedisSetting.Host, RedisSetting.Port),
		Password: RedisSetting.Password,    // no password set
		DB:       viper.GetInt("redis.db"), // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
	})
	err = redisHelper.Ping(ctx).Err()
	if err != nil {
		fmt.Println("ping redis fail!\n", err)
		return err
	}
	return
}
func GetRedisHelper() *RedisHelper {
	return redisHelper
}
func RedisClose() {
	_ = redisHelper.Close()
}
