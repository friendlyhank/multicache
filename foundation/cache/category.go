package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetCategory(categoryid int64)(category *db.Category,err error){
	category = &db.Category{}
	if err = CategoryCache.Get(category,categoryid);err != nil{
		category = nil
	}
	return
}
