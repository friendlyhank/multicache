package multicache

import (
	"context"
	"github.com/friendlyhank/groupcache"
)

//MultiCache- 二层缓存
// 第一层: 基于GroupCache的分布式客户端缓存
// 第二层: 基于Redis的中央缓存
// 第三层: MySQL的存储缓存
type MultiCache struct{
	localCache *LocalCache
	rdsCache *RedisCache
	getter     Getter //db 数据源
}

// Getter -
type Getter interface {
	Get(ds interface{}, args ...interface{}) error
}

// GetterFunc -
type GetterFunc func(ds interface{}, args ...interface{})  error

// Get -
func (f GetterFunc) Get(ds interface{}, args ...interface{}) error {
	return f(ds, args...)
}

func (r *MultiCache)Set(){}

func (r *MultiCache)Get(){}

func (r *MultiCache)Remove(){
}

func MakeMultiCache(getter Getter)*MultiCache{
	if getter == nil{
		panic("multiCache nil Getter")
	}
	multiCache := &MultiCache{getter:getter}

	//make rediscache
	multiCache.rdsCache = MakeRedisCache(GetterFunc(func(ds interface{},args ...interface{})error{
		if err := getter.Get(ds,args);err != nil{
			return err
		}
		return nil
	}))

	//make localcache
	multiCache.localCache = MakeLocalCache(groupcache.GetterFunc(func(ctx context.Context,key string,dest groupcache.Sink)error{
		return nil
	}))

	return multiCache
}
