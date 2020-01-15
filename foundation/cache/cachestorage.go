package cache

import (
	"errors"
	"github.com/friendlyhank/multicache"
	"github.com/friendlyhank/multicache/foundation/db"
)

// 默认的缓存参数
var (
	ErrNotFound = errors.New("NotFound")
)

var(
	ActorCache *multicache.MultiCache
)

// InitCacheStorage - 初始化缓存
func Init() {
	redisexpired := 5
	localexpired := 2
	ActorCache = multicache.MakeMultiCache("actor",redisexpired,localexpired,10000, multicache.GetterFunc(getByI64))
}

// getByI64 - 单一的主键查询
func getByI64(ds interface{}, args ...interface{}) (err error) {
	var has bool
	has, err = db.Engine().ID(args[0].(int64)).Get(ds)
	if err != nil {
		ds = nil
	} else if !has {
		ds = nil
		err = ErrNotFound
	}

	return
}