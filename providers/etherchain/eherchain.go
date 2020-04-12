package etherchain

import (
	"encoding/json"
	"github.com/OsoianMarcel/egpp/common"
	"strconv"
)

type apiJsonResponse struct {
	Fast     string `json:"fast"`
	Fastest  string `json:"fastest"`
	SafeLow  string `json:"safeLow"`
	Standard string `json:"standard"`
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

	standard, err := strconv.ParseFloat(r.Standard, 16)
	if err != nil {
		return common.GasPrice{}, err
	}
	safeLow, err := strconv.ParseFloat(r.SafeLow, 16)
	if err != nil {
		return common.GasPrice{}, err
	}
	fast, err := strconv.ParseFloat(r.Fast, 16)
	if err != nil {
		return common.GasPrice{}, err
	}
	fastest, err := strconv.ParseFloat(r.Fastest, 16)
	if err != nil {
		return common.GasPrice{}, err
	}

	return common.GasPrice{
		Standard: uint16(standard),
		SafeLow:  uint16(safeLow),
		Fast:     uint16(fast),
		Fastest:  uint16(fastest),
		Provider: p.GetName(),
	}, nil
}
