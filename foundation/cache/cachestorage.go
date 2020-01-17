package cache

import (
	"errors"
	"github.com/friendlyhank/multicache"
	"github.com/friendlyhank/multicache/foundation/db"
	"github.com/xormplus/core"
)

// 默认的缓存参数
var (
	ErrNotFound = errors.New("NotFound")
)

var(
	ActorCache *multicache.MultiCache
	AddressCache *multicache.MultiCache
	CategoryCache *multicache.MultiCache
	CityCache  *multicache.MultiCache
	CountryCache *multicache.MultiCache
	CustomerCache *multicache.MultiCache
	FilmCache *multicache.MultiCache
	FilmActorCache *multicache.MultiCache
	FilmCategoryCache *multicache.MultiCache
	FilmTextCache *multicache.MultiCache
	InventoryCache *multicache.MultiCache
	LanguageCache *multicache.MultiCache
	PaymentCache *multicache.MultiCache
	RentalCache *multicache.MultiCache
	StaffCache *multicache.MultiCache
	StoreCache *multicache.MultiCache

)

// InitCacheStorage - 初始化缓存
func Init() {
	redisexpired := 5 * 60
	localexpired := 3 *60
	var cacheBytes int64 = 10 * 1024 * 1024
	ActorCache = multicache.MakeMultiCache("actor",
		redisexpired, localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	AddressCache = multicache.MakeMultiCache("address",
		redisexpired,localexpired,cacheBytes,
		multicache.GetterFunc(getByI64))
	CategoryCache = multicache.MakeMultiCache("category",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	CityCache = multicache.MakeMultiCache("city",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	CountryCache = multicache.MakeMultiCache("country",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	CustomerCache = multicache.MakeMultiCache("customer",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	FilmCache = multicache.MakeMultiCache("film",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	FilmActorCache = multicache.MakeMultiCache("filmactor",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByMulti))
	FilmCategoryCache = multicache.MakeMultiCache("filmcategory",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByMulti))
	FilmTextCache = multicache.MakeMultiCache("filmtext",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	InventoryCache = multicache.MakeMultiCache("inventory",
		redisexpired,localexpired,
		cacheBytes, multicache.GetterFunc(getByI64))
	LanguageCache = multicache.MakeMultiCache("language",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	PaymentCache = multicache.MakeMultiCache("payment",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	RentalCache = multicache.MakeMultiCache("rental",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	StaffCache = multicache.MakeMultiCache("staff",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
	StoreCache = multicache.MakeMultiCache("store",
		redisexpired,localexpired,
		cacheBytes,
		multicache.GetterFunc(getByI64))
}

// getByI64 - 单一的主键查询
func getByI64(ds interface{}, args ...interface{}) (err error) {
	var has bool
	has, err = db.Engine().ID(args[0].(int64)).Get(ds)
	if err != nil {
		ds = nil
	}else if !has{
		ds = nil
		err = ErrNotFound
	}
	return
}

// getByMulti - 多主键查询
func getByMulti(ds interface{}, args ...interface{}) (err error) {
	var has bool
	has, err = db.Engine().Id(core.NewPK(args...)).Get(ds)
	if err != nil {
		ds = nil
	} else if !has {
		ds = nil
		err = ErrNotFound
	}
	return
}