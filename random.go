package balance

import (
	"errors"
	"math/rand"
)

type randomBalance struct {
	currentIndex int
	address []string
	length int
}

func (r *randomBalance) GetAddress() string {
	r.currentIndex = rand.Intn(r.length)

	return r.address[r.currentIndex]
}

func NewRandomBalance(address []string) (*randomBalance, error) {
	data := &randomBalance{
		currentIndex: 0,
		address: address,
		length: len(address),
	}

	if data.length <= 0 {
		return nil, errors.New("list is empty")
	}

	return data, nil
}
