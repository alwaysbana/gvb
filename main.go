package main

import (
	"gvb/core"
	"gvb/global"
)

func main() {
	//读取
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("日志初始化警告!")
	global.Log.Error("日志初始化错误!")
	global.Log.Info("日志初始化信息!")
	//global.DB = core.InitGorm()
	//fmt.Println("ok")

	//fmt.Println(global.DB)
	//fmt.Println("ok")
}
