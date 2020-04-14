package ethergasstation

import (
	"encoding/json"
	"github.com/OsoianMarcel/egpp/common"
)

type apiJsonResponse struct {
	Fast    float32 `json:"fast"`
	Fastest float32 `json:"fastest"`
	SafeLow float32 `json:"safeLow"`
	Average float32 `json:"average"`
}

type Provider struct{}

func NewProvider() Provider {
	return Provider{}
}

func (p Provider) GetName() string {
	return "Ether Gas Station"
}

func (p Provider) Request() (common.GasPrice, error) {
	res, err := common.GetRequest("https://ethgasstation.info/json/ethgasAPI.json")
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
		Standard: uint16(r.Average) * 10,
		SafeLow:  uint16(r.SafeLow) * 10,
		Fast:     uint16(r.Fast) * 10,
		Fastest:  uint16(r.Fastest) * 10,
		Provider: p.GetName(),
	}, nil
}
