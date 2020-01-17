package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetCity(cityid int64)(city *db.City,err error){
	city = &db.City{}
	if err = CityCache.Get(city,cityid);err != nil{
		city = nil
	}
	return
}
