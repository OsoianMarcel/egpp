package egpp

import (
	"errors"

	"github.com/OsoianMarcel/egpp/common"
)

// The method requests the gas price from specified providers in order and returns the first successful result
func GetGasPriceWithFallback(providers []common.Provider) (common.GasPrice, error) {
	for _, provider := range providers {
		gasPrice, err := provider.Request()
		if err == nil {
			return gasPrice, nil
		}
	}

	return common.GasPrice{}, errors.New("all providers failed")
}

// The method requests gas prices from all specified providers in parallel and returns a result with average values
func GetGasPriceAverage(providers []common.Provider) (common.GasPrice, error) {
	results := common.BatchRequests(providers)

	gasPrices := make([]common.GasPrice, 0, len(results))
	for _, gpr := range results {
		if gpr.Err == nil {
			gasPrices = append(gasPrices, gpr.GasPrice)
		}
	}

	if len(gasPrices) == 0 {
		return common.GasPrice{}, errors.New("all providers failed")
	}

	return common.AverageGasPrice(gasPrices), nil
}
