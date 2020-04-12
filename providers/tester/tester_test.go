package tester_test

import (
	"github.com/OsoianMarcel/egpp/common"
	"github.com/OsoianMarcel/egpp/providers/tester"
	"testing"
)

func TestProvider_GetName(t *testing.T) {
	p := tester.NewProvider(false, common.GasPrice{})

	if p.GetName() != "Tester" {
		t.Error("wrong name")
	}
}

func TestProvider_Request_Failure(t *testing.T) {
	p := tester.NewProvider(false, common.GasPrice{})

	_, err := p.Request()
	if err == nil {
		t.Error("request should always return an error")
	}
}

func TestProvider_Request_Success(t *testing.T) {
	providedGasPrice := common.GasPrice{
		Provider: "None",
		Standard: 3,
		SafeLow:  1,
		Fast:     5,
		Fastest:  12,
	}
	p := tester.NewProvider(true, providedGasPrice)

	gasPrice, err := p.Request()
	if err != nil {
		t.Error("request should not return an error")
		return
	}

	if gasPrice != providedGasPrice {
		t.Error("provided gas price and returned gas price are different")
	}
}
