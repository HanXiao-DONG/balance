package test

import (
	"github.com/HanXiao-DONG/balance"
	"testing"
)

func TestRandom(t *testing.T) {
	random, error := balance.NewRandomBalance([]string{
		"1111111",
		"2222222",
		"3333333",
	})

	if error != nil {
		t.Errorf("excepted:%v", error)
		return
	}

	for i := 0; i <= 10; i++ {
		t.Log(random.GetAddress())
	}
}
