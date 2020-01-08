package multicache

import (
	"context"
	"encoding/json"
	"fmt"
	rds "github.com/friendlyhank/goredis"
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
func key(prefix string,args ...interface{})string{
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

//getArgsByKey-
func getArgsByKey(key string)[]interface{}{
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
func (m *MultiCache)Set(val interface{}, key string) error {
	return nil
}

func (m *MultiCache)SetExpired(val interface{}, key string,expired int) error {
	m.localCache.Set(val,key)
	m.rdsCache.SetExpired(val,key,expired)
	return nil
}

//Get-
func (m *MultiCache)Get(val interface{},args ...interface{})error{
	var err error
	//开关
	if m.localCache != nil{
		err = m.localCache.Get(val,args)
	}else if m.rdsCache != nil{
		err = m.rdsCache.Get(val,args)
	}

	return err
}

//Remove-
func (m *MultiCache)Remove(){
}

func MakeMultiCache(name string,redisexpired,localexpired int,cacheBytes int64,getter Getter)*MultiCache{
	if getter == nil{
		panic("multiCache nil Getter")
	}
	multiCache := &MultiCache{getter:getter}

	//make rediscache
	multiCache.rdsCache = MakeRedisCache(name,redisexpired,rds.GetRedisDefault(),
		GetterFunc(func(ds interface{},args ...interface{})error{
		if err := getter.Get(ds,args);err != nil{
			return err
		}
		return nil
	}))

	//make localcache
	multiCache.localCache = MakeLocalCache(name,cacheBytes,int64(localexpired),
		groupcache.GetterFunc(func(ctx context.Context,key string,dest groupcache.Sink)error{

		args := getArgsByKey(key)
		var val interface{}
		if err := multiCache.rdsCache.Get(val,args);err != nil{
			return err
		}

		v,_ := json.Marshal(val)
		dest.SetBytes(v)

		return nil
	}))

	return multiCache
}
