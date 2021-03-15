package egpp_test

import (
	"testing"

	"github.com/OsoianMarcel/egpp"
	"github.com/OsoianMarcel/egpp/common"
	"github.com/OsoianMarcel/egpp/providers/ethergasstation"
	"github.com/OsoianMarcel/egpp/providers/tester"
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
			Standard: 30,
			SafeLow:  10,
			Fast:     50,
			Fastest:  200,
		}),
		tester.NewProvider(true, common.GasPrice{
			Provider: "1",
			Standard: 20,
			SafeLow:  10,
			Fast:     0,
			Fastest:  150,
		}),
	}

	avgGasPrice, err := egpp.GetGasPriceAverage(providers)
	if err != nil {
		t.Error(err)
		return
	}

	if avgGasPrice.Provider != "Average Gas Price" {
		t.Error("average gas price provider is wrong")
		return
	}

	if avgGasPrice.Standard != 25 ||
		avgGasPrice.SafeLow != 10 ||
		avgGasPrice.Fast != 50 ||
		avgGasPrice.Fastest != 175 {
		t.Error("average gas price values are wrong")
	}
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
