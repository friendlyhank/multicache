package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetStore(storeid int64)(store *db.Store,err error){
	store = &db.Store{}
	if err = StoreCache.Get(store,storeid);err != nil{
		store = nil
	}
	return
}
