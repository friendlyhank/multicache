package multicache

//MultiCache- 二层缓存
// 第一层: 基于GroupCache的分布式客户端缓存
// 第二层: 基于Redis的中央缓存
// 第三层: MySQL的存储缓存
type MultiCache struct{
	localCache *LocalCache
	rdsCache *RedisCache
	getter     Getter //数据源(这里默认DB)
}

// Getter -
type Getter interface {
	Get(ds interface{}, args ...interface{}) (err error)
}

// GetterFunc -
type GetterFunc func(ds interface{}, args ...interface{}) (err error)

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
		panic("nil Getter")
	}
	multiCache := &MultiCache{getter:getter}

	//make rediscache
	multiCache.rdsCache = MakeRedisCache(getter)

	//make localcache
	multiCache.localCache = MakeLocalCache(getter)

	return multiCache
}
