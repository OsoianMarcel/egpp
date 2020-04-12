package egpp_test

import (
	"fmt"
	"github.com/OsoianMarcel/egpp"
	"github.com/OsoianMarcel/egpp/common"
	"github.com/OsoianMarcel/egpp/providers/ethergasstation"
	"github.com/OsoianMarcel/egpp/providers/tester"
	"testing"
)

func TestGetGasPriceWithFallback_FallbackSuccess(t *testing.T) {
	providers := []common.Provider{
		tester.NewProvider(false, common.GasPrice{}),
		ethergasstation.NewProvider(),
	}

	gasPrice, err := egpp.GetGasPriceWithFallback(providers)
	if err != nil {
		t.Error(err)
		return
	}

	if gasPrice.Provider != "Ether Gas Station" {
		t.Error("fetched provider name should be different")
	}
}

func TestGetGasPriceWithFallback_AllFails(t *testing.T) {
	providers := []common.Provider{
		tester.NewProvider(false, common.GasPrice{}),
		tester.NewProvider(false, common.GasPrice{}),
	}

	_, err := egpp.GetGasPriceWithFallback(providers)
	if err == nil {
		t.Error("the method should return an error")
	}
}

func TestGetGasPriceAverage_Success(t *testing.T) {
	providers := []common.Provider{
		tester.NewProvider(true, common.GasPrice{
			Provider: "1",
			Standard: 3,
			SafeLow:  1,
			Fast:     5,
			Fastest:  12,
		}),
		tester.NewProvider(true, common.GasPrice{
			Provider: "1",
			Standard: 2,
			SafeLow:  1,
			Fast:     0,
			Fastest:  10,
		}),
	}

	avgGasPrice, err := egpp.GetGasPriceAverage(providers)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("\n%+v\n", avgGasPrice)
}

func TestGetGasPriceAverage_AllFail(t *testing.T) {
	providers := []common.Provider{
		tester.NewProvider(false, common.GasPrice{}),
		tester.NewProvider(false, common.GasPrice{}),
	}

	_, err := egpp.GetGasPriceAverage(providers)
	if err == nil {
		t.Error("get gas price average should return an error")
	}
}
