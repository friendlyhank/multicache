package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetStaff(staffid int64)(staff *db.Staff,err error){
	staff = &db.Staff{}
	if err = StaffCache.Get(staff,staffid);err != nil{
		staff = nil
	}
	return
}
