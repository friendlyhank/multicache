package test

import (
	"fmt"
	_ "github.com/friendlyhank/multicache/foundation"
	"github.com/friendlyhank/multicache/foundation/cache"
	"github.com/friendlyhank/multicache/foundation/db"
	"math/rand"
	"testing"
	"time"
)

var actorids []int64


func setUp(){
	fmt.Println("data source init")
	db.Engine().Table("actor").Cols("actor_id").Find(&actorids)
}

func TestMain(m *testing.M){
	setUp()
	m.Run()
}

func TestMultiCache(t *testing.T){
	for{
		RandomGetActior(t)
		time.Sleep(3 * time.Second)
	}
}

//RandomGetActior -
func RandomGetActior(t *testing.T){
	randId := rand.Intn(len(actorids)-1)
	actor,err := cache.GetActor(int64(randId))
	if err != nil{
		fmt.Println("get actor|Err|",err.Error())
	}
	fmt.Println("|get actor|actor|",actor)
}
