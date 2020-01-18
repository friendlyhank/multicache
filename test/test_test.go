package test

import (
	"fmt"
	"github.com/friendlyhank/multicache"
	_ "github.com/friendlyhank/multicache/foundation"
	"github.com/friendlyhank/multicache/foundation/cache"
	"github.com/friendlyhank/multicache/foundation/db"
	"math/rand"
	"testing"
	"time"
)

var(
	actorids []int64
	addressids []int64
	categoryids []int64
	cityids []int64
	countryids []int64
	customerids []int64
	filmids []int64
	filmtids []int64
	inventoryids []int64
	languageids []int64
	paymentids []int64
	rentalids []int64
	staffids []int64
	storeids []int64

	fimActorpk []*FimActorPk
	filmcategorypk []*FilmCategoryPk
)

//FimActorPk-
type FimActorPk struct{
	ActorId    int       `xorm:"not null pk SMALLINT(5)"`
	FilmId     int       `xorm:"not null pk index SMALLINT(5)"`
}

//FimActorPk-
type FilmCategoryPk struct{
	FilmId     int       `xorm:"not null pk SMALLINT(5)"`
	CategoryId int       `xorm:"not null pk index TINYINT(3)"`
}

func setUp(){
	fmt.Println("data source init")

	db.Engine().Table("actor").Cols("actor_id").Find(&actorids)
	db.Engine().Table("address").Cols("address_id").Find(&addressids)
	db.Engine().Table("category").Cols("category_id").Find(&categoryids)
	db.Engine().Table("city").Cols("city_id").Find(&cityids)
	db.Engine().Table("country").Cols("country_id").Find(&countryids)
	db.Engine().Table("customer").Cols("customer_id").Find(&customerids)
	db.Engine().Table("film").Cols("film_id").Find(&filmids)
	db.Engine().Table("film_text").Cols("film_id").Find(&filmtids)
	db.Engine().Table("inventory").Cols("inventory_id").Find(&inventoryids)
	db.Engine().Table("language").Cols("language_id").Find(&languageids)
	db.Engine().Table("payment").Cols("payment_id").Find(&paymentids)
	db.Engine().Table("rental").Cols("rental_id").Find(&rentalids)
	db.Engine().Table("staff").Cols("staff_id").Find(&staffids)
	db.Engine().Table("store").Cols("store_id").Find(&storeids)

	db.Engine().Table("film_actor").Cols("actor_id","film_id").Find(&fimActorpk)
	db.Engine().Table("film_category").Cols("film_id","category_id").Find(&filmcategorypk)
}

func TestMain(m *testing.M){
	setUp()
	m.Run()
}

func TestMultiCache(t *testing.T){
	for{
		randTypeId := rand.Intn(15)
		SwitchGetCache(randTypeId)

		//time.Sleep(time.Duration(randTypeId) * time.Second)
		time.Sleep(1 * time.Second)
	}
}

func SwitchGetCache(randTypeId int){
	switch randTypeId {
	case 0:
		RandomGetActior()
	case 1:
		RandomGetAddress()
	case 2:
		RandomGetCategory()
	case 3:
		RandomGetCity()
	case 4:
		RandomGetCountry()
	case 5:
		RandomGetCustomer()
	case 6:
		RandomGetFilm()
	case 7:
		RandomGetFilmActor()
	case 8:
		RandomGetFilmCategory()
	case 9:
		RandomGetFilmText()
	case 10:
		RandomGetInventory()
	case 11:
		RandomGetLanguage()
	case 12:
		RandomGetPayment()
	case 13:
		RandomGetRental()
	case 14:
		RandomGetStaff()
	case 15:
		RandomGetStore()
	}
}

//RandomGetActior -
func RandomGetActior(){
	randId := rand.Intn(len(actorids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	actor,err := cache.GetActor(int64(actorids[randId]))
	timeCost(start,"actor",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get actor|Err|",err.Error())
	}
	fmt.Println("|get actor|",actor)
}

//RandomGetAddress -
func RandomGetAddress(){
	randId := rand.Intn(len(addressids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	address,err := cache.GetAddress(int64(addressids[randId]))
	timeCost(start,"address",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get address|Err|",err.Error())
	}
	fmt.Println("|get address|",address)
}

//RandomGetCategory -
func RandomGetCategory(){
	randId := rand.Intn(len(categoryids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	category,err := cache.GetCategory(int64(categoryids[randId]))
	timeCost(start,"category",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get category|Err|",err.Error())
	}
	fmt.Println("|get category|",category)
}

//RandomGetCity -
func RandomGetCity(){
	randId := rand.Intn(len(cityids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	city,err := cache.GetCity(int64(cityids[randId]))
	timeCost(start,"city",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get city|Err|",err.Error())
	}
	fmt.Println("|get city|",city)
}

//RandomGetCountry -
func RandomGetCountry(){
	randId := rand.Intn(len(countryids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	country,err := cache.GetCountry(int64(countryids[randId]))
	timeCost(start,"country",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get country|Err|",err.Error())
	}
	fmt.Println("|get country|",country)
}

//RandomGetCustomer -
func RandomGetCustomer(){
	randId := rand.Intn(len(customerids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	customer,err := cache.GetCustomer(int64(customerids[randId]))
	timeCost(start,"customer",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get customer|Err|",err.Error())
	}
	fmt.Println("|get customer|",customer)
}

//RandomGetFilm -
func RandomGetFilm(){
	randId := rand.Intn(len(filmids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	film,err := cache.GetFilm(int64(filmids[randId]))
	timeCost(start,"film",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get film|Err|",err.Error())
	}
	fmt.Println("|get film|",film)
}

//RandomGetFilmActor -
func RandomGetFilmActor(){
	randId := rand.Intn(len(fimActorpk)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	filmActor,err := cache.GetFilmActor(int64(fimActorpk[randId].ActorId),int64(fimActorpk[randId].FilmId))
	timeCost(start,"filmActor",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get filmActor|Err|",err.Error())
	}
	fmt.Println("|get filmActor|",filmActor)
}

//RandomGetFilmCategory -
func RandomGetFilmCategory(){
	randId := rand.Intn(len(filmcategorypk)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	filmCategory,err := cache.GetFilmCategory(int64(filmcategorypk[randId].FilmId),int64(filmcategorypk[randId].CategoryId))
	timeCost(start,"filmCategory",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get filmCategory|Err|",err.Error())
	}
	fmt.Println("|get filmCategory|",filmCategory)
}

//RandomGetFilmText -
func RandomGetFilmText(){
	randId := rand.Intn(len(filmtids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	filmText,err := cache.GetFilmText(int64(filmtids[randId]))
	timeCost(start,"filmText",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get filmText|Err|",err.Error())
	}
	fmt.Println("|get filmText|",filmText)
}

//RandomGetInventory -
func RandomGetInventory(){
	randId := rand.Intn(len(inventoryids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	inventory,err := cache.GetInventory(int64(inventoryids[randId]))
	timeCost(start,"inventory",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get actor|Err|",err.Error())
	}
	fmt.Println("|get inventory|",inventory)
}

//RandomGetLanguage -
func RandomGetLanguage(){
	randId := rand.Intn(len(languageids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	language,err := cache.GetLanguage(int64(languageids[randId]))
	timeCost(start,"language",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get actor|Err|",err.Error())
	}
	fmt.Println("|get language|",language)
}

//RandomGetPayment -
func RandomGetPayment(){
	randId := rand.Intn(len(paymentids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	payment,err := cache.GetPayment(int64(paymentids[randId]))
	timeCost(start,"payment",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get payment|Err|",err.Error())
	}
	fmt.Println("|get payment|",payment)
}

//RandomGetRental -
func RandomGetRental(){
	randId := rand.Intn(len(rentalids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	lental,err := cache.GetRental(int64(rentalids[randId]))
	timeCost(start,"lental",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get lental|Err|",err.Error())
	}
	fmt.Println("|get lental|",lental)
}

//RandomGetStaff -
func RandomGetStaff(){
	randId := rand.Intn(len(staffids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	staff,err := cache.GetStaff(int64(staffids[randId]))
	timeCost(start,"staff",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get staff|Err|",err.Error())
	}
	fmt.Println("|get staff|",staff)
}

//RandomGetStore -
func RandomGetStore(){
	randId := rand.Intn(len(storeids)-1)

	beginStats := *multicache.GetStats()
	start := time.Now()
	store,err := cache.GetStore(int64(storeids[randId]))
	timeCost(start,"store",beginStats,*multicache.GetStats())

	if err != nil{
		fmt.Println("get store|Err|",err.Error())
	}
	fmt.Println("|get store|",store)
}

//@brief：耗时统计函数
func timeCost(start time.Time,dbName string,beginStats multicache.Stats,Endstats multicache.Stats){
	tc:=time.Since(start)
	var cacheType string
	if beginStats.CacheHits != Endstats.CacheHits{
		cacheType = "local"
	}
	if beginStats.RedisLoads != Endstats.RedisLoads{
		cacheType = "redis"
	}
	if beginStats.LocalLoads != Endstats.LocalLoads{
		cacheType = "db"
	}
	if cacheType != ""{
		fmt.Printf("cache：%v;dbname：%v;time cost = %v\n",cacheType,dbName,tc)
	}else{
		fmt.Printf("something worng; time cost = %v\n",tc)
	}

}



