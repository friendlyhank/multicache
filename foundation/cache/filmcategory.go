package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetFilmCategory(filmid,categoryid int64)(filmCategory *db.FilmCategory,err error){
	filmCategory = &db.FilmCategory{}
	if err = FilmCategoryCache.Get(filmCategory,filmid,categoryid);err != nil{
		filmCategory = nil
	}
	return
}
