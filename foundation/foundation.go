package foundation

import (
	"github.com/friendlyhank/multicache/foundation/cache"
	"github.com/friendlyhank/multicache/foundation/goredis"
	"github.com/friendlyhank/multicache/foundation/db"
)

func init(){
	//数据库初始化
	db.Init()
	//redis初始化
	rds.Init()
	//cache初始化
	cache.Init()
}
