package cache

import "github.com/friendlyhank/multicache/foundation/db"

func GetInventory(inventoryid int64)(inventory *db.Inventory,err error){
	inventory = &db.Inventory{}
	if err = InventoryCache.Get(inventory,inventoryid);err != nil{
		inventory = nil
	}
	return
}
