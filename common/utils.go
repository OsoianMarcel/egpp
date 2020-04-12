package common

import (
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func BatchRequests(providers []Provider) []GasPriceResult {
	n := len(providers)
	c := make(chan GasPriceResult, n)
	results := make([]GasPriceResult, 0, n)

	for _, provider := range providers {
		go func(p Provider, ic chan GasPriceResult) {
			gasPrice, err := p.Request()
			ic <- GasPriceResult{GasPrice: gasPrice, Err: err}
		}(provider, c)
	}

	for i := 0; i < n; i++ {
		results = append(results, <-c)
	}

	return results
}

func AverageGasPrice(gasPrices []GasPrice) GasPrice {
	avgGp := GasPrice{}

	var safeLowN, standardN, fastN, fastestN uint16

	for _, gp := range gasPrices {
		avgGp.SafeLow += gp.SafeLow
		avgGp.Standard += gp.Standard
		avgGp.Fast += gp.Fast
		avgGp.Fastest += gp.Fastest

		if gp.SafeLow > 0 {
			safeLowN++
		}
		if gp.Standard > 0 {
			standardN++
		}
		if gp.Fast > 0 {
			fastN++
		}
		if gp.Fastest > 0 {
			fastestN++
		}
	}

	avgGp.Provider = "Average Gas Price"
	if safeLowN > 0 {
		avgGp.SafeLow = avgGp.SafeLow / safeLowN
	}
	if standardN > 0 {
		avgGp.Standard = avgGp.Standard / standardN
	}
	if fastN > 0 {
		avgGp.Fast = avgGp.Fast / fastN
	}
	if fastestN > 0 {
		avgGp.Fastest = avgGp.Fastest / fastestN
	}

	return avgGp
}
