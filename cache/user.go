package cache

import (
	"GinDemo/global"
	"GinDemo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// token的过期时长
const TokenDuration = time.Minute * 5

// cache的名字
func getTokenCacheName(articleId uint64) string {
	return "article_" + strconv.FormatUint(articleId, 10)
}

// SetAndGetToken Login (redis)
func SetAndGetToken(context *gin.Context, value string, expiration time.Duration) (token string) {
	token = "token_demo"
	err := global.GetRedisHelper().Set(context, token, value, TokenDuration).Err()
	if err != nil {
		fmt.Printf("redis set fail!\nerr:", err)
		return
	}
	return
}

func UserRank(context *gin.Context, user *models.LoginUserBo) {
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
