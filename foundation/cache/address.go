package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetAddress(addressid int64)(address *db.Address,err error){
	address = &db.Address{}
	if err = AddressCache.Get(address,addressid);err != nil{
		address = nil
	}
	return
}