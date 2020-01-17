package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetFilmText(filmid int64)(filmText *db.FilmText,err error){
	filmText = &db.FilmText{}
	if err = FilmTextCache.Get(filmText,filmid);err != nil{
		filmText = nil
	}
	return
}
