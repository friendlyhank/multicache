package multicache

import "github.com/friendlyhank/groupcache"

//LocalCache -
type LocalCache struct{
	prefix string
	GroupCache *groupcache.Group
	//数据源
}

func (r *LocalCache)Set(){

}

func (r *LocalCache)Get(){

}

func (r *LocalCache)Remove(){
}

func MakeLocalCache(getter Getter)*LocalCache{
	r := &LocalCache{}
	return r
}
