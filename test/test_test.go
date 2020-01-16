package test

import (
	_ "github.com/friendlyhank/multicache/foundation"
	"github.com/friendlyhank/multicache/foundation/cache"
	"github.com/friendlyhank/multicache/foundation/db"
	"math/rand"
	"testing"
	"time"
)

var actorids []int64

func TestMultiCache(t *testing.T){
	for{
		RandomGetActior(t)
		time.Sleep(3 * time.Second)
	}
}

//RandomGetActior -
func RandomGetActior(t *testing.T){
	if actorids == nil{
		db.Engine().Table("actor").Cols("actor_id").Find(&actorids)
	}
	randId := rand.Intn(len(actorids)-1)
	actor,err := cache.GetActor(int64(randId))
	if err != nil{
		t.Errorf("get actor|Err|%v|",err)
	}
	t.Logf("|get actor|id|%v|",actor.ActorId)
}
