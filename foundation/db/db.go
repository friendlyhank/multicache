package db

import(
	"github.com/astaxie/beego/logs"
	"github.com/xormplus/xorm"
)

var(
	mallEngine *xorm.Engine
)

func Init(){
	logs.Debug("|foundation|init|db|Init")

}
