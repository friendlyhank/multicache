package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetActor(actorid int64)(actor *db.Actor,err error){
	actor = &db.Actor{}
	if err = ActorCache.Get(actor,actorid);err != nil{
		actor = nil
	}
	return
}
