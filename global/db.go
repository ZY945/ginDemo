package global

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// MySQLInit 初始化数据库
func MySQLInit() (err error) {
	//"user:password@tcp(127.0.0.1:3306)/dao?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DatabaseSetting.UserName, DatabaseSetting.Password, DatabaseSetting.Host, DatabaseSetting.Port, DatabaseSetting.DBName)
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	DB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	return
}
func MysqlClose() {
	_ = DB.Close()
}
