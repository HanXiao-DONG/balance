package balance

import (
	"errors"
)

type pollBalance struct {
	currentIndex int
	address []string
	length int
}

func (p *pollBalance) GetAddress() string {
	addr := p.address[p.currentIndex]

	p.currentIndex = (p.currentIndex + 1) % p.length

	return addr
}

func NewPollBalance(address []string) (*pollBalance, error) {
	data := &pollBalance{
		currentIndex: 0,
		address: address,
		length: len(address),
	}

	if data.length <= 0 {
		return nil, errors.New("address is empty")
	}

	return data, nil
}
