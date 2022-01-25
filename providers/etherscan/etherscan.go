package etherscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/OsoianMarcel/egpp/common"
)

type apiResult struct {
	SafeGasPrice    string
	ProposeGasPrice string
}

type apiResponse struct {
	Status string    `json:"status"`
	Result apiResult `json:"result"`
}

type Provider struct {
	apiKey string
}

func NewProvider(apiKey string) Provider {
	return Provider{apiKey}
}

func (p Provider) GetName() string {
	return "Etherscan"
}

func (p Provider) Request() (common.GasPrice, error) {
	providerUrl := fmt.Sprintf("https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=%s", p.apiKey)
	resp, err := common.GetRequest(providerUrl)
	if err != nil {
		return common.GasPrice{}, err
	}

	return p.parseApiResponse(resp)
}

func (p Provider) parseApiResponse(resp []byte) (common.GasPrice, error) {
	var r apiResponse
	err := json.Unmarshal(resp, &r)
	if err != nil {
		return common.GasPrice{}, err
	}

	if r.Status != "1" {
		return common.GasPrice{}, errors.New("etherscan status must be 1")
	}

	proposeGasPrice, err := strconv.ParseUint(r.Result.ProposeGasPrice, 10, 16)
	if err != nil {
		return common.GasPrice{}, err
	}
	safeLow, err := strconv.ParseUint(r.Result.SafeGasPrice, 10, 16)
	if err != nil {
		return common.GasPrice{}, err
	}

	return common.GasPrice{
		Standard: uint16(safeLow) * 1e2,
		Fastest:  uint16(proposeGasPrice) * 1e2,
		Provider: p.GetName(),
	}, nil
}
