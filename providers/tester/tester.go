package tester

import (
	"errors"
	"github.com/OsoianMarcel/egpp/common"
)

type Provider struct {
	success  bool
	gasPrice common.GasPrice
}

func NewProvider(success bool, gasPrice common.GasPrice) Provider {
	return Provider{success, gasPrice}
}

func (p Provider) GetName() string {
	return "Tester"
}

func (p Provider) Request() (common.GasPrice, error) {
	if p.success {
		return p.gasPrice, nil
	}

	return p.gasPrice, errors.New("permanent error")
}
