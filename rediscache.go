package multicache

import rds "github.com/friendlyhank/goredis"

type RedisCache struct{
	prefix string
	RedisSource *rds.RedisSource
	getter     Getter //db 数据源
}

func (r *RedisCache)Set(){

}

func (r *RedisCache)Get(){

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
