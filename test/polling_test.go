package test

import (
	"github.com/HanXiao-DONG/balance"
	"testing"
)

func TestPolling(t *testing.T) {
	polling, error := balance.NewPollBalance([]string{
		"1111111",
		"2222222",
		"3333333",
	})
	if error != nil {
		t.Errorf("excepted:%v", error)
		return
	}

	for i := 0; i <= 10; i++ {
		t.Log(polling.GetAddress())
	}
}
