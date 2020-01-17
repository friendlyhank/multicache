package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetPayment(paymentid int64)(payment *db.Payment,err error){
	payment = &db.Payment{}
	if err = PaymentCache.Get(payment,paymentid);err != nil{
		payment = nil
	}
	return
}
