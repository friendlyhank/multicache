package db

import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var(
	mallEngine *xorm.Engine
)

func Init(){
	logs.Debug("|foundation|init|db|Init")
	var(
		err error
	)

	dbSource := beego.AppConfig.DefaultString("dbsource","root:123456@tcp(127.0.0.1:3306)/sakila?charset=utf8mb4&loc=Local&interpolateParams=true")
	if mallEngine,err = xorm.NewEngine("mysql",dbSource);err != nil{
		logs.Error("Engine Init Err:%v", err)
		panic(err)
	}

	mallEngine.ShowSQL(true)
	mallEngine.ShowExecTime(true)
	mallEngine.SetLogger(xorm.NewSimpleLogger(beego.BeeLogger))
	mallEngine.Logger().SetLevel(core.LOG_INFO)
}

//Engine -主数据库
func Engine()*xorm.Engine{
	return mallEngine
}
