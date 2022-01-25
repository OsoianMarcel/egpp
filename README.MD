# EGPP - Ethereum Gas Price Provider
A Golang library used to fetch **Ethereum Gas Prices** from different gas price providers.

[![Build Status](https://app.travis-ci.com/OsoianMarcel/egpp.svg?branch=master)](https://app.travis-ci.com/OsoianMarcel/egpp)
[![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/OsoianMarcel/egpp)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/OsoianMarcel/egpp/blob/master/LICENSE)

## Basic example
```go
package main

import (
	"fmt"
	"log"

	"github.com/OsoianMarcel/egpp"
	"github.com/OsoianMarcel/egpp/common"
	"github.com/OsoianMarcel/egpp/providers/etherchain"
	"github.com/OsoianMarcel/egpp/providers/ethergasstation"
	"github.com/OsoianMarcel/egpp/providers/etherscan"
)

func main() {
    // Gas price providers
	providers := []common.Provider{
		ethergasstation.NewProvider(),
		etherchain.NewProvider(),
		etherscan.NewProvider(""),
	}

    // Requests the gas price values from providers in order and returns the first successful result
	gasPriceStruct, err := egpp.GetGasPriceWithFallback(providers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", gasPriceStruct)
    // {Provider:Ether Gas Station Standard:9800 SafeLow:9200 Fast:14300 Fastest:14300}
}
```

## Contribute

Contributions to the package are always welcome!

* Report any bugs or issues you find on the [issue tracker].
* You can grab the source code at the package's [Git repository].

## License

All contents of this package are licensed under the [MIT license].

[issue tracker]: https://github.com/OsoianMarcel/egpp/issues
[Git repository]: https://github.com/OsoianMarcel/egpp
[MIT license]: LICENSE