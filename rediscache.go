package multicache

import rds "github.com/friendlyhank/goredis"

type RedisCache struct{
	prefix string
	expired int
	RedisSource *rds.RedisSource
	getter     Getter //db 数据源
}

//Get-
func (r *RedisCache)Get(val interface{},args ...interface{})error{
	key :=key(r.prefix,args)

	//
	if _,err:=r.RedisSource.GetJSON(key,val);err == nil{
		//找到了
		return nil
	}
	//
	if r.getter == nil{
		panic("NoLocalGetterInRedis")
		return nil
	}
	if err := r.getter.Get(val,args...);err != nil{
		//没有找到
		return err
	}

	//找到了设置到rds
	err := r.SetExpired(val,key,r.expired)

	return err
}

//SetExpired-
func (r *RedisCache)SetExpired(val interface{}, key string, expired int)error{
	err := r.RedisSource.SetJSON(key,val,expired)
	return err
}


func (r *RedisCache)Remove(){

}

func MakeRedisCache(getter Getter)*RedisCache{
	if getter == nil{
		panic("redisCache nil Getter")
	}
	r := &RedisCache{getter:getter}
	return r
}
