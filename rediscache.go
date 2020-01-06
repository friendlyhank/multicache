package multicache

import rds "github.com/friendlyhank/goredis"

type RedisCache struct{
	prefix string
	RedisSource *rds.RedisSource
	getter     Getter //数据源(这里默认DB)
}

func (r *RedisCache)Set(){

}

func (r *RedisCache)Get(){

}

func (r *RedisCache)Remove(){

}

func MakeRedisCache(getter Getter)*RedisCache{

}
