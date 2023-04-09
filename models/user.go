package models

import (
	"database/sql"
)

type User struct {
	Id         int            `db:"id"`
	UserName   sql.NullString `db:"username" json:"username"`
	PassWord   sql.NullString `db:"password" json:"password"`
	Name       sql.NullString `db:"name" json:"name"`
	Age        sql.NullInt64  `db:"age" json:"age"`
	Sex        sql.NullString `db:"sex" json:"sex"`
	Address    sql.NullString `db:"address" json:"address"`
	Phone      sql.NullString `db:"phone" json:"phone"`
	CreateTime sql.NullTime   `db:"create_time" json:"createTime"`
}

type UserVo struct {
	Id         int
	UserName   string
	Name       string
	Age        int64
	Sex        string
	Address    string
	Phone      string
	CreateTime string
}
type LoginUserBo struct {
	UserName string
	PassWord string
}

type AddUserBo struct {
	UserName string
	PassWord string
	Name     string
	Age      int64
	Sex      string
	Address  string
	Phone    string
	//CreateTime time.Time `default:"3"` //为0时
}

type UpdateUserBo struct {
	Id      int
	Age     int64
	Address string
	Phone   string
}
