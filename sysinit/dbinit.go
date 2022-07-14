package sysinit


import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "zs403_mbook_copy/models" // 初始化 models 的 init 函数
)

//调用方式
//dbinit() 或 dbinit("w") 或 dbinit("default") //初始化主库
//dbinit("w","r")	//同时初始化主库和从库
//dbinit("w")
func dbinit(aliases ...string) {
	isDev := ("dev" == beego.AppConfig.String("runmode"))  // 读取配置文件的值

	if len(aliases) > 0 {
		for _,alias := range aliases {
			registDatabase(alias)
			// 主库 自动建表
			if "w" == alias {
				// 参数1 : 表的别名，默认是 "default",
				// 参数2 : 报错是否强行执行下一条 sql
				// 参数3 : 是否打印输出调试信息 true 是； false 否
				orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registDatabase("w")
		// 参数1 : 表的别名，默认是 "default",
		// 参数2 : 报错是否强行执行下一条 sql
		// 参数3 : 是否打印输出调试信息 true 是； false 否
		orm.RunSyncdb("default", false, isDev)
	}
}

func registDatabase(alias string)  {
	if len(alias) == 0 {
		return
	}
	//连接名称
	dbAlias := alias
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
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

	test := beego.AppConfig.String("db_w_username_1111")  // 读取 db.conf 的配置文件
	fmt.Println(test)

	// 设置数据库连接参数。使用数据库驱动程序
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)
}