package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetRental(rentalid int64)(lental *db.Rental,err error){
	lental = &db.Rental{}
	if err = RentalCache.Get(lental,rentalid);err != nil{
		lental = nil
	}
	return
}
