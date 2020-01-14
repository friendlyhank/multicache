package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetActor(actor_id int64)(actor *db.Actor,err error){
	if err = ActorCache.Get(actor,actor);err != nil{
		actor = nil
	}
	return
}
