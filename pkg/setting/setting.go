package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

// NewSetting Init 也可以使用ini.Load()去读取ini文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config") // 指定配置文件(不需要带后缀)
	vp.AddConfigPath("config") //指定配置文件类型
	vp.SetConfigType("yaml")   // 指定查找配置文件的路径
	err := vp.ReadInConfig()   // 读取配置信息
	if err != nil {
		return nil, err // 读取配置信息失败
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	s := &Setting{vp}
	return s, nil
}

// 读取指定的一段
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
