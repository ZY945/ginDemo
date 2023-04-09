package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本,如果爆红,把包删除即可,重新import
	"github.com/spf13/viper"
)

// 声明一个全局变量
var (
	RedisDb *redis.Client
	ctx     = context.Background()
)

// RedisInit Init 初始化连接
func RedisInit() (err error) {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", RedisSetting.Host, RedisSetting.Port),
		Password: RedisSetting.Password,    // no password set
		DB:       viper.GetInt("redis.db"), // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	err = RedisDb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("ping redis fail!\n", err)
		return err
	}
	return
}

func RedisClose() {
	_ = RedisDb.Close()
}
