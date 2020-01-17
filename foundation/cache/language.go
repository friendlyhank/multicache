package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetLanguage(languageid int64)(language *db.Language,err error){
	language = &db.Language{}
	if err = LanguageCache.Get(language,languageid);err != nil{
		language = nil
	}
	return
}
