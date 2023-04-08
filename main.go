package main

import (
	"GinAndSqlx/dao/mysql"
	"GinAndSqlx/routers"
	"GinAndSqlx/setting"
	"fmt"
)

func main() {
	//1.加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}
	//3.初始化mysql
	if err := mysql.SqlxInitMySQL(); err != nil {

		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	router := routers.InitRouter()
	router.Run(":8080")
}
