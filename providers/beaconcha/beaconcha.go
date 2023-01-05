package beaconcha

import (
	"encoding/json"

	"github.com/OsoianMarcel/egpp/common"
)

type dataJsonResponse struct {
	Rapid    uint64 `json:"rapid"`
	Fast     uint64 `json:"fast"`
	Standard uint64 `json:"standard"`
	Slow     uint64 `json:"slow"`
}

type rootJsonResponse struct {
	Code uint8            `json:"code"`
	Data dataJsonResponse `json:"data"`
}

type Provider struct{}

func NewProvider() Provider {
	return Provider{}
}

func (p Provider) GetName() string {
	return "Beaconcha"
}

func (p Provider) Request() (common.GasPrice, error) {
	res, err := common.GetRequest("https://beaconcha.in/api/v1/execution/gasnow")
	if err != nil {
		return common.GasPrice{}, err
	}

	return p.parseApiResponse(res)
}

func (p Provider) parseApiResponse(resp []byte) (common.GasPrice, error) {
	var r rootJsonResponse
	err := json.Unmarshal(resp, &r)
	if err != nil {
		return common.GasPrice{}, err
	}

	return common.GasPrice{
		Standard: uint16(r.Data.Standard / 1e7),
		SafeLow:  uint16(r.Data.Slow / 1e7),
		Fast:     uint16(r.Data.Fast / 1e7),
		Fastest:  uint16(r.Data.Rapid / 1e7),
		Provider: p.GetName(),
	}, nil
}
