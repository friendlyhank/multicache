package multicache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/friendlyhank/groupcache"
)

//LocalCache -
type LocalCache struct{
	prefix string
	groupCache *groupcache.Group
	getter     Getter //redis 数据源
}

type GetterContext struct{
	val interface{}
	args []interface{}
}

func (l *LocalCache)Set(val interface{},key string)error{
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	slink := groupcache.AllocatingByteSliceSink(&b)
	l.groupCache.Set(key,slink)

	return nil
}

//SetExpired- 暂不支持失效
func (l *LocalCache)SetExpired(val interface{},key string, expired int) error{
	return nil
}

func (l *LocalCache)Get(val interface{}, args ...interface{})error{
	key := genkey(l.prefix,args...)
	err := l.groupCache.Get(nil, key,groupcache.AllJsonSink(val))

	fmt.Println(l.groupCache.Stats.Gets)
	fmt.Println(l.groupCache.Stats.CacheHits)
	fmt.Println(l.groupCache.Stats.PeerLoads)
	fmt.Println(l.groupCache.Stats.PeerErrors)
	fmt.Println(l.groupCache.Stats.LocalLoads)
	fmt.Println(l.groupCache.Stats.LocalLoadErrs)
	fmt.Println(l.groupCache.Stats.ServerRequests)
	return err
}

func (l *LocalCache)Remove(){
}

func MakeLocalCache(name string,cacheBytes int64,expired int64,getter Getter)*LocalCache{
	if getter == nil{
		panic("localCache nil Getter")
	}
	l := &LocalCache{
		prefix:name,
		getter:getter,
	}
	l.groupCache = groupcache.NewGroupExt(name,cacheBytes,expired,
		groupcache.GetterFunc(func(ctx context.Context,key string,dest groupcache.Sink)error{
			args := getArgsByKey(key)
			getterContext := dest.(*groupcache.JsonSink)
			if err :=getter.Get(getterContext.Dst,args...);err != nil{
				return err
			}
			b,_ := json.Marshal(getterContext.Dst)
			dest.SetBytes(b)
			return nil
	}))
	return l
}
