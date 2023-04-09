package gorm

import "time"

// User gorm 类型名就是表名
type User struct {
	Id         int       `gorm:"primary_key;auto_increment" json:"id" db:"id"`
	UserName   string    `gorm:"type:varchar(20);not null" db:"username" json:"username"`
	PassWord   string    `gorm:"type:varchar(500);not null" json:"password"`
	Name       string    `gorm:"type:varchar(20)" json:"name"`
	Age        int64     `gorm:"type:int;not null" db:"age" json:"age" `
	Sex        string    `gorm:"type:varchar(20)" db:"sex" json:"sex"`
	Address    string    `gorm:"type:varchar(20)" db:"address" json:"address"`
	Phone      string    `gorm:"type:varchar(20)" db:"phone" json:"phone"`
	CreateTime time.Time `gorm:"type:datetime;not null;" json:"create_time" db:"create_time"`
}
