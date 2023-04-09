package global

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var DB *sqlx.DB
var GormDB *gorm.DB

// MySQLInit 初始化数据库
func MySQLInit() (err error) {
	//"user:password@tcp(127.0.0.1:3306)/dao?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", DatabaseSetting.UserName, DatabaseSetting.Password, DatabaseSetting.Host, DatabaseSetting.Port, DatabaseSetting.DBName)
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

// GormInit 初始化gorm
func GormInit() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", DatabaseSetting.UserName, DatabaseSetting.Password, DatabaseSetting.Host, DatabaseSetting.Port, DatabaseSetting.DBName)
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	sqlDB, _ := GormDB.DB()
	DB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	DB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return
}
