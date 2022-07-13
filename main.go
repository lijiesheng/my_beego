package main

import (
	_ "zs403_mbook_copy/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

