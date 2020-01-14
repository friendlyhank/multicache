package test

import (
	"github.com/friendlyhank/multicache/foundation/cache"
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
