package main

import (
	"GinDemo/global"
	"GinDemo/routers"
	"fmt"
)

func main() {
	//1.加载配置
	if err := global.SetupSetting(); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}
	//2.初始化mysql
	if err := global.MySQLInit(); err != nil {

		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer global.MysqlClose()

	if err := global.GormInit(); err != nil {
		fmt.Printf("init gorm failed, err:%v\n", err)
		return
	}
	//2.初始化redis
	if err := global.RedisInit(); err != nil {

		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer global.RedisClose()

	//雪花算法
	if err := global.SnowFlakeInit(); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	router := routers.InitRouter()

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("service run fail!")
		return
	}
}
