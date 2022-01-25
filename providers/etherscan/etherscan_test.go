package etherscan_test

import (
	"testing"

	"github.com/OsoianMarcel/egpp/providers/etherscan"
)

func TestProvider_GetName(t *testing.T) {
	p := etherscan.NewProvider("")

	if p.GetName() != "Etherscan" {
		t.Error("wrong name")
	}
}

func TestProvider_Request(t *testing.T) {
	p := etherscan.NewProvider("")

	gasPrice, err := p.Request()
	if err != nil {
		t.Error(err)
		return
	}

	if gasPrice.Standard == 0 || gasPrice.Fastest == 0 {
		t.Error("all gas price properties should be greater than zero")
	}

	if gasPrice.Provider != p.GetName() {
		t.Error("different provider name")
	}
}
