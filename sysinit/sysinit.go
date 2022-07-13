package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
)

func registDatabase(alias string)  {
	if len(alias) == 0 {
		return
	}
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")

	fmt.Println(dbName)
	fmt.Println(dbUser)
	fmt.Println(dbPwd)
	fmt.Println(dbHost)
	fmt.Println(dbPort)
}