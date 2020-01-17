package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetFilmActor(actorid,filmid int64)(filmActor *db.FilmActor,err error){
	filmActor = &db.FilmActor{}
	if err = FilmActorCache.Get(filmActor,actorid,filmid);err != nil{
		filmActor = nil
	}
	return
}
