package global

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis"
)

// 声明一个全局变量
var (
	RedisDb *redis.Client
)

// RedisInit Init 初始化连接
func RedisInit() (err error) {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", RedisSetting.Host, RedisSetting.Port),
		Password: RedisSetting.Password,    // no password set
		DB:       viper.GetInt("redis.db"), // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = RedisDb.Ping().Result()

	return err
}

func RedisClose() {
	_ = RedisDb.Close()
}
