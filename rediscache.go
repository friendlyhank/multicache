package multicache

import (
	rds "github.com/friendlyhank/multicache/foundation/goredis"
	"github.com/gomodule/redigo/redis"
	pfc "github.com/niean/goperfcounter"
)

type RedisCache struct{
	prefix string
	expired int
	RedisSource *rds.RedisSource
	getter     Getter //db 数据源
}

//Get-
func (r *RedisCache)Get(val interface{},args ...interface{})error{

	var err error
	key :=genkey(r.prefix,args...)

	//
	if _,err =r.RedisSource.GetJSON(key,val);err == nil{
		//找到了
		allstats.RedisLoads.Add(1)
		pfc.Gauge("multicache.redisloads.count", int64(allstats.RedisLoads))
		return nil
	}
	if err != nil && err != redis.ErrNil{
		allstats.RedisLoadErrs.Add(1)
		pfc.Gauge("multicache.redisloaderrs.count", int64(allstats.RedisLoadErrs))
	}
	//
	if r.getter == nil{
		panic("NoLocalGetterInRedis")
		return nil
	}
	if err := r.getter.Get(val,args...);err != nil{
		allstats.LocalLoadErrs.Add(1)
		pfc.Gauge("multicache.localloaderrs.count", int64(allstats.LocalLoadErrs))
		//没有找到
		return err
	}
	allstats.LocalLoads.Add(1)
	pfc.Gauge("multicache.localloads.count", int64(allstats.LocalLoads))

	//找到了设置到rds
	err = r.SetExpired(val,key,r.expired)

	return err
}

//SetExpired-
func (r *RedisCache)SetExpired(val interface{}, key string, expired int)error{
	err := r.RedisSource.SetJSON(key,val,expired)
	return err
}


func (r *RedisCache)Remove(){

}

func MakeRedisCache(name string,expired int,source *rds.RedisSource,getter Getter)*RedisCache{
	if getter == nil{
		panic("redisCache nil Getter")
	}
	r := &RedisCache{
		prefix:name,
		getter:getter,
		RedisSource:source,
		expired:expired,
	}
	r.getter = GetterFunc(func(ds interface{}, args ...interface{})  error{
		var err error
		if err  = getter.Get(ds,args...);err != nil{
			return err
		}
		return nil
	})
	return r
}
