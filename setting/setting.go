package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件(不需要带后缀)
	viper.SetConfigType("yaml")   //指定配置文件类型
	viper.AddConfigPath(".")      // 指定查找配置文件的路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {               // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}
