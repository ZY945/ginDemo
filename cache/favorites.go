package cache

import (
	"GinDemo/global"
	"GinDemo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func addRedis() {

}

func FavoritesRank(context *gin.Context, user *models.LoginUserBo) {
	zsetUserKey := "user_rank"
	score := global.GetRedisHelper().ZScore(context, zsetUserKey, user.UserName).Val()
	fmt.Println(score)
	if score != -1 {
		global.GetRedisHelper().ZIncrBy(context, zsetUserKey, 1, user.UserName)
	} else {
		_, err := global.GetRedisHelper().ZAdd(context, zsetUserKey, &redis.Z{
			Score:  1,
			Member: user,
		}).Result()
		if err != nil {
			fmt.Printf("rank zadd init fail:\n%s", err)
			return
		}
	}

}
