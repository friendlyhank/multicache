package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetCustomer(customerid int64)(customer *db.Customer,err error){
	customer = &db.Customer{}
	if err = CustomerCache.Get(customer,customerid);err != nil{
		customer = nil
	}
	return
}
