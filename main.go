package main

import (
	"github.com/astaxie/beego"
	_ "zs403_mbook_copy/routers" // 初始化 routers 的 init 函数
	_ "zs403_mbook_copy/sysinit" // 这里要先初始化 sysinit 的 init 函数
)

func main() {
	beego.Run()
}


