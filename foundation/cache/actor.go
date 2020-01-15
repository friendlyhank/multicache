package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetActor(actor_id int64)(actor *db.Actor,err error){
	actor = &db.Actor{}
	if err = ActorCache.Get(actor,actor_id);err != nil{
		actor = nil
	}
	return
}
