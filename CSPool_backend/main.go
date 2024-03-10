package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mysqlmodule "main/dao/mysql"
	redismodule "main/dao/redis"
	"main/router"
)

func main() {
	//初始化MySQL连接
	if err := mysqlmodule.Init(); err != nil {
		fmt.Printf("MySQL start failed, err:%v\n", err)
	}
	defer mysqlmodule.Sdb.Close()
	//初始化Redis连接
	if err := redismodule.Init(); err != nil {
		fmt.Printf("Redis start failed, err:%v\n", err)
	}
	defer redismodule.Rdb.Close()
	//注册路由
	r := router.Init()
	//启动服务
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Server start failed, err:%v\n", err)
		return
	}
}
