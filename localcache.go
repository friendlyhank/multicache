package multicache

import (
	"context"
	"encoding/json"
	"github.com/friendlyhank/groupcache"
)

//LocalCache -
type LocalCache struct{
	prefix string
	groupCache *groupcache.Group
	getter     Getter //redis 数据源
}

type GetterContext struct{
	Val interface{}
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

	var b []byte
	l.groupCache.Get(nil,key,groupcache.AllocatingByteSliceSink(&b))
	err := json.Unmarshal(b,val)

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
			var getterContext = &GetterContext{}
			if err :=getter.Get(getterContext,args...);err != nil{
				return err
			}
			b,_ := json.Marshal(getterContext.Val)
			dest.SetBytes(b)
			return nil
	}))
	return l
}
