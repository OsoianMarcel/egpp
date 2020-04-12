package common

type GasPrice struct {
	Provider string `json:"provider"`
	Standard uint16 `json:"standard"`
	SafeLow  uint16 `json:"safe_low"`
	Fast     uint16 `json:"fast"`
	Fastest  uint16 `json:"fastest"`
}

type GasPriceResult struct {
	GasPrice GasPrice `json:"gas_price"`
	Err      error    `json:"err"`
}
