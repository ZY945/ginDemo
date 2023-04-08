package models

import "database/sql"

//type User struct {
//	Id         int    `db:"id"`
//	Name       string `db:"name"`
//	Age        int    `db:"age"`
//	Sex        string `db:"sex"`
//	Address    string `db:"address"`
//	Phone      string `db:"phone"`
//	CreateTime string `db:"create_time"`
//}

type User struct {
	Id         int            `db:"id"`
	Name       sql.NullString `db:"name"`
	Age        sql.NullInt64  `db:"age"`
	Sex        sql.NullString `db:"sex"`
	Address    sql.NullString `db:"address"`
	Phone      sql.NullString `db:"phone"`
	CreateTime sql.NullString `db:"create_time"`
}

type UserVo struct {
	Id         int
	Name       string
	Age        int64
	Sex        string
	Address    string
	Phone      string
	CreateTime string
}

type AddUserBo struct {
	Name       string
	Age        int64
	Sex        string
	Address    string
	Phone      string
	CreateTime string
}

type UpdateUserBo struct {
	Id      int
	Age     int64
	Address string
	Phone   string
}
