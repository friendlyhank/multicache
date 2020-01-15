package test

import (
	"github.com/friendlyhank/multicache/foundation/cache"
	"github.com/friendlyhank/multicache/foundation/db"
	"testing"
	_ "github.com/friendlyhank/multicache/foundation"
)

func TestMultiCache(t *testing.T){
	actor,err := cache.GetActor(1)
	if err != nil{
		t.Errorf("%v",err)
	}
	t.Logf("%v",actor)
}

func TestDBCache(t *testing.T){
	var actor = &db.Actor{}
	has,err := db.Engine().ID(1).Get(actor)
	if err != nil || !has{
		actor = nil
		t.Errorf("%v",err)
	}

	t.Logf("%v",actor)
}
