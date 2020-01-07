package multicache

import "github.com/friendlyhank/groupcache"

//LocalCache -
type LocalCache struct{
	prefix string
	GroupCache *groupcache.Group
	getter     Getter //redis 数据源
}

type localCacheContxt struct{
	val  interface{}
	args []interface{}
}

func (r *LocalCache)Set(){

}

func (r *LocalCache)Get(){

}

func (r *LocalCache)Remove(){
}

func MakeLocalCache(getter groupcache.Getter)*LocalCache{
	if getter == nil{
		panic("localCache nil Getter")
	}
	r := &LocalCache{}
	return r
}
