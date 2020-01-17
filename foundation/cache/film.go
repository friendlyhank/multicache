package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetFilm(filmid int64)(film *db.Film,err error){
	film = &db.Film{}
	if err = FilmCache.Get(film,filmid);err != nil{
		film = nil
	}
	return
}
