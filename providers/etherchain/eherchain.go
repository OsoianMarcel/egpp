package etherchain

import (
	"encoding/json"

	"github.com/OsoianMarcel/egpp/common"
)

type apiJsonResponse struct {
	Fast     float32 `json:"fast"`
	Fastest  float32 `json:"fastest"`
	SafeLow  float32 `json:"safeLow"`
	Standard float32 `json:"standard"`
}

type Provider struct{}

func NewProvider() Provider {
	return Provider{}
}

func (p Provider) GetName() string {
	return "Etherchain"
}

func (p Provider) Request() (common.GasPrice, error) {
	res, err := common.GetRequest("https://www.etherchain.org/api/gasPriceOracle")
	if err != nil {
		return common.GasPrice{}, err
	}

	return p.parseApiResponse(res)
}

func (p Provider) parseApiResponse(resp []byte) (common.GasPrice, error) {
	var r apiJsonResponse
	err := json.Unmarshal(resp, &r)
	if err != nil {
		return common.GasPrice{}, err
	}

	return common.GasPrice{
		Standard: uint16(r.Standard * 1e2),
		SafeLow:  uint16(r.SafeLow * 1e2),
		Fast:     uint16(r.Fast * 1e2),
		Fastest:  uint16(r.Fastest * 1e2),
		Provider: p.GetName(),
	}, nil
}
