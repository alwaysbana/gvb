package main

import (
	"fmt"
	"gvb/core"
	"gvb/global"
)

func main() {
	//读取
	core.InitConf()
	global.DB = core.InitGorm()
	//fmt.Println("ok")

	fmt.Println(global.DB)
	//fmt.Println("ok")
}
