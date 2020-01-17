package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetCountry(cityid int64)(country *db.Country,err error){
	country = &db.Country{}
	if err = CountryCache.Get(country,cityid);err != nil{
		country = nil
	}
	return
}
