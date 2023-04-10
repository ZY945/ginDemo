package gorm

import "time"

type Favorites struct {
	Id           int        `gorm:"primary_key;auto_increment" db:"id" json:"id"`
	Name         string     `gorm:"type:varchar(20)" db:"name" json:"name" `
	Url          string     `gorm:"type:varchar(255)" db:"url" json:"url"`
	Introduction string     `gorm:"type:varchar(255)" db:"introduction" json:"introduction"`
	Directory    string     `gorm:"type:varchar(20)" db:"directory" json:"directory"` // 多个url在一个Directory
	CreatedTime  *time.Time `gorm:"type:TIMESTAMP" db:"created_time" json:"created_time" `
	UpdateTime   *time.Time `gorm:"type:TIMESTAMP" db:"update_time" json:"update_time"`
	Deleted      int        `gorm:"type:int" db:"deleted" json:"deleted"` // 0未删除,1删除
}

func (*Favorites) TableName() string {
	return "favorites"
}

type AddFavoritesBo struct {
	Name         string
	Url          string
	Introduction string
	Directory    string
	CreatedTime  *time.Time
}

type SelectFavoritesBo struct {
	Id           int
	Name         string
	Url          string
	Introduction string
	Directory    string
	CreatedTime  *time.Time
}

func GetFavoritesByBo(bo *AddFavoritesBo) (favorites *Favorites) {
	return &Favorites{
		Name:         bo.Name,
		Url:          bo.Url,
		Introduction: bo.Introduction,
		Directory:    bo.Directory,
		CreatedTime:  bo.CreatedTime,
	}
}
