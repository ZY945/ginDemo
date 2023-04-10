package dao

import (
	"GinDemo/global"
	"GinDemo/models/gorm"
	"fmt"
	"time"
)

func GetFavoritesById(id int) (favorites *gorm.Favorites) {
	err := global.GormDB.Model(&gorm.Favorites{}).Where("id = ? and deleted = ?", id, 0).First(&favorites).Error
	if err != nil {
		fmt.Printf("select by id fail:\n%s", err)
	}
	return
}

func GetFavoritesList(bo *gorm.SelectFavoritesBo) (favorites *[]gorm.Favorites, length int64) {
	favorites = new([]gorm.Favorites)
	tx := global.GormDB.Model(&gorm.Favorites{})
	if bo.Id > 0 {
		tx.Where("id = ?", bo.Id)
	}
	if bo.Name != "" {
		tx.Where("name = ?", bo.Name)
	}
	if bo.Url != "" {
		tx.Where("url = ?", bo.Url)
	}
	if bo.Introduction != "" {
		tx.Where("introduction = ?", bo.Introduction)
	}
	if bo.Directory != "" {
		tx.Where("directory = ?", bo.Directory)
	}
	result := tx.Where("deleted = ?", 0).Find(&favorites)
	err := result.Error
	length = result.RowsAffected
	if err != nil {
		fmt.Printf("select by id fail:\n%s", err)
	}
	return
}
func AddFavorites(bo *gorm.AddFavoritesBo) {
	//设置时区
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Printf("The time zone error:\n%s", err)
	//	return
	//}
	time_string := time.Now().Format("2006-01-02 15:04:05")
	location, _ := time.Parse("2006-01-02 15:04:05", time_string)
	bo.CreatedTime = &location
	favorites := gorm.GetFavoritesByBo(bo)
	result := global.GormDB.Model(&favorites).Create(&favorites)

	if result.Error != nil {
		fmt.Printf("AddFavorites fail:\n %s", result.Error)
		return
	}

}
func DelFavorites(id int) {
	err := global.GormDB.Model(&gorm.Favorites{}).Where("id = ?", id).Update("deleted", 1).Error
	if err != nil {
		fmt.Printf("del fail:\n%s", err)
		return
	}
}
