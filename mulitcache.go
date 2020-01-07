package multicache

import (
	"context"
	"fmt"
	"github.com/friendlyhank/groupcache"
	"strings"
)


const (
	cachePrefix    = "__cache_"
	cachePrefixLen = len(cachePrefix)
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


func cacheName(prefix string)string{
	return fmt.Sprintf("%v%v", cachePrefix, prefix)
}

//key-
func (r *MultiCache)key(prefix string,args ...interface{})string{
	if len(args) == 0{
		return ""
	}
	key := fmt.Sprintf("%v", args[0])
	if len(args) > 1 {
		for _, v := range args[1:] {
			key = fmt.Sprintf("%v_%v", key, v)
		}
	}
	return cacheName(prefix) + "-"+ key
}

//splitkey-
func splitkey(key string)[]interface{}{
	n := strings.Index(key, "-")
	if n < cachePrefixLen {
		return nil
	}

	shortkey := key[n+1:]
	splits := strings.Split(shortkey,"_")

	var args []interface{}
	for _,s := range splits{
		args = append(args,s)
	}

	return args
}

//Set -
func (r *MultiCache)Set(){}

//Get-
func (r *MultiCache)Get(){}

//Remove-
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
