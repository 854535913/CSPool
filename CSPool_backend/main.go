package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/appconfig"
	mysqlmodule "main/dao/mysql"
	redismodule "main/dao/redis"
	"main/router"
)

func main() {
	// 加载配置
	if err := appconfig.Init("./appconfig/config_local.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	//初始化MySQL连接
	if err := mysqlmodule.Init(appconfig.Conf.MySQLConfig); err != nil {
		fmt.Printf("MySQL start failed, err:%v\n", err)
	}
	defer mysqlmodule.Close()
	//初始化Redis连接
	if err := redismodule.Init(appconfig.Conf.RedisConfig); err != nil {
		fmt.Printf("Redis start failed, err:%v\n", err)
	}
	defer redismodule.Close()
	//注册路由
	r := router.Init()
	//启动服务
	if err := r.Run(fmt.Sprintf(":%d", appconfig.Conf.Port)); err != nil {
		fmt.Printf("Server start failed, err:%v\n", err)
		return
	}
}
