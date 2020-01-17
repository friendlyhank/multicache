package multicache

import (
	"fmt"
	rds "github.com/friendlyhank/multicache/foundation/goredis"
	"strconv"
	"strings"
	"sync/atomic"
)


const (
	cachePrefix    = "__cache_"
	cachePrefixLen = len(cachePrefix)
)

// An AtomicInt is an int64 to be accessed atomically.
type AtomicInt int64

// Add atomically adds n to i.
func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

// Get atomically gets the value of i.
func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

func (i *AtomicInt) String() string {
	return strconv.FormatInt(i.Get(), 10)
}

// Stats are per-group statistics.
type Stats struct {
	Gets           AtomicInt //请求数
	CacheHits      AtomicInt //本地缓存
	RedisLoads     AtomicInt //远端缓存
	RedisLoadErrs AtomicInt  //远端错误
	LocalLoads     AtomicInt //Db请求
	LocalLoadErrs  AtomicInt //Db错误
	ServerRequests AtomicInt //网络请求次数
}

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

//genkey-
func genkey(prefix string,args ...interface{})string{
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
		id,_ := strconv.ParseInt(s, 10, 64)
		args = append(args,id)
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
		err = m.localCache.Get(val,args...)
	}else if m.rdsCache != nil{
		err = m.rdsCache.Get(val,args...)
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

	//make rediscache 数据库源
	multiCache.rdsCache = MakeRedisCache(name,redisexpired,rds.GetRedisDefault(), getter)

	//make localcache redis源
	multiCache.localCache = MakeLocalCache(name,cacheBytes,int64(localexpired),GetterFunc(multiCache.rdsCache.Get))

	return multiCache
}
