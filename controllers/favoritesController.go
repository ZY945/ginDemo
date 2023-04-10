package controllers

import (
	"GinDemo/dao"
	"GinDemo/models/gorm"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func AddFavorites(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var add *gorm.AddFavoritesBo
	err := json.Unmarshal(data, &add)
	if err != nil {
		log.Print("json error:\n", err)
		return
	}
	fmt.Println(add)
	dao.AddFavorites(add)

}
func DelFavorites(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	dao.DelFavorites(id)
}
func GetFavoritesById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	favorites := dao.GetFavoritesById(id)
	if favorites.Id == 0 {
		ctx.JSON(200, gin.H{
			"message":   "没有该对象",
			"favorites": nil,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message":   "查询成功",
			"favorites": favorites,
		})
	}
}
func GetFavoritesList(ctx *gin.Context) {
	idStr := ctx.Query("id")
	nameStr := ctx.Query("name")
	urlStr := ctx.Query("url")
	introductionStr := ctx.Query("introduction")
	directoryStr := ctx.Query("directory")
	id, _ := strconv.Atoi(idStr)
	bo := &gorm.SelectFavoritesBo{
		Id:           id,
		Name:         nameStr,
		Url:          urlStr,
		Introduction: introductionStr,
		Directory:    directoryStr,
	}
	fmt.Print(bo)
	favorites, length := dao.GetFavoritesList(bo)
	if favorites == nil {
		ctx.JSON(200, gin.H{
			"length":    0,
			"message":   "列表为空",
			"favorites": nil,
		})
	} else {
		ctx.JSON(200, gin.H{
			"length":    length,
			"message":   "查询成功",
			"favorites": favorites,
		})
	}
}
