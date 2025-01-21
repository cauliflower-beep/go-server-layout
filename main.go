package main

import (
	"app-server/pkg/logger"
	"app-server/pkg/snowflake"
	"app-server/router"
	"app-server/settings"
	"fmt"
)

func main() {

	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 初始化日志
	if err := logger.Init(settings.GetConf().LogConfig, settings.GetConf().Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 初始化数据库连接 Mysql | MongoDB | Redis

	// 初始化雪花算法
	if err := snowflake.Init(settings.GetConf().StartTime, settings.GetConf().MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 业务模块初始化 如自定义的定时任务等

	// 注册信号量 实现服务优雅启停 todo

	// 注册路由
	r := router.SetupRouter(settings.GetConf().Mode)
	// 启动服务
	err := r.Run(fmt.Sprintf("%s:%d", settings.GetConf().Addr, settings.GetConf().Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
