package beaconcha_test

import (
	"testing"

	"github.com/OsoianMarcel/egpp/providers/beaconcha"
)

func TestProvider_GetName(t *testing.T) {
	p := beaconcha.NewProvider()

	if p.GetName() != "Beaconcha" {
		t.Error("wrong name")
	}
}

func TestProvider_Request(t *testing.T) {
	p := beaconcha.NewProvider()

	gasPrice, err := p.Request()
	if err != nil {
		t.Error(err)
		return
	}

	if gasPrice.Standard == 0 || gasPrice.SafeLow == 0 || gasPrice.Fast == 0 || gasPrice.Fastest == 0 {
		t.Error("all gas price properties should be greater than zero")
	}

	if gasPrice.Provider != p.GetName() {
		t.Error("different provider name")
	}
}
