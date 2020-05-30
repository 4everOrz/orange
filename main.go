package main

import (
	"github.com/alecthomas/log4go"
	"orange/common/config"
	_ "orange/common/log"
	"orange/db"
	_ "orange/db"
	"orange/migrate"
	"orange/runtime"
	"orange/service/webService"
)

func main() {

	//数据库
	database:=db.New()
	database.Init()
	if err:=database.Connect();err!=nil{
		log4go.Error("数据库连接异常:",err.Error())
		return
	}
	defer database.Close()
	//数据库自动迁移
	 migrate:=migrate.New(database.DB())
	 migrate.AutoMigrate()
     //创建运行时
     runtime:=runtime.New(database.DB())
     runtime.Init()
	//web服务
	listenPort,_:=config.GetString("server","Listen")
     web:=webService.New(listenPort,runtime)
     web.Start()
}
