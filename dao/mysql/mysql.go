package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

/*
	sqlx框架
*/
// 初始化数据库
func SqlxInitMySQL() (err error) {
	//"user:password@tcp(127.0.0.1:3306)/dao?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetInt("mysql.port"), viper.GetString("mysql.dbname"))
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}
func Close() {
	_ = db.Close()
}
