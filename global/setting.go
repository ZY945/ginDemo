package global

import (
	"GinAndSqlx/pkg/setting"
	"time"
)

// 定义全局变量
var (
	ServerSetting    *ServerSettingS
	DatabaseSetting  *DatabaseSettingS
	RedisSetting     *RedisSettingS
	SnowFlakeSetting *SnowFlakeSettingS
)

// 服务器配置
type ServerSettingS struct {
	Name         string
	RunMode      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 数据库配置
type DatabaseSettingS struct {
	Host         string
	Port         int
	UserName     string
	Password     string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxOpenConns int
	MaxIdleConns int
}

// redis配置
type RedisSettingS struct {
	Host     string
	Port     int
	Password string
	PoolSize int
	DB       int
}

// redis配置
type SnowFlakeSettingS struct {
	StartTime string `mapstructure:"start_time"`
	MachineId int    `mapstructure:"machine_id"`
}

// 读取配置到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("mysql", &DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("app", &ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("redis", &RedisSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("snowflake", &SnowFlakeSetting)
	if err != nil {
		return err
	}

	return nil
}
