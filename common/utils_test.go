package common_test

import (
	"bytes"
	"github.com/OsoianMarcel/egpp/common"
	"github.com/OsoianMarcel/egpp/providers/tester"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest_Error(t *testing.T) {
	_, err := common.GetRequest("http://localhost:-41")
	if err == nil {
		t.Error("the function should return an error")
	}
}

func TestGetRequest_Success(t *testing.T) {
	expectedResponse := []byte("Ok")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(expectedResponse)
		if err != nil {
			t.Error(err)
		}
	}))
	defer ts.Close()

	res, err := common.GetRequest(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(res, expectedResponse) {
		t.Errorf("expected response: %s, result: %s", expectedResponse, res)
	}
}

func TestAverageGasPrice(t *testing.T) {
	avgGasPrice := common.AverageGasPrice([]common.GasPrice{
		common.GasPrice{
			Provider: "1",
			Standard: 2,
			SafeLow:  0,
			Fast:     5,
			Fastest:  12,
		},
		common.GasPrice{
			Provider: "2",
			Standard: 3,
			SafeLow:  0,
			Fast:     7,
			Fastest:  11,
		},
		common.GasPrice{
			Provider: "3",
			Standard: 3,
			SafeLow:  0,
			Fast:     5,
			Fastest:  0,
		},
	})

	if avgGasPrice.Provider != "Average Gas Price" {
		t.Error("wrong provider name")
	}

	if avgGasPrice.Standard != 2 {
		t.Error("wrong standard")
	}
	if avgGasPrice.SafeLow != 0 {
		t.Error("wrong safelow")
	}
	if avgGasPrice.Fast != 5 {
		t.Error("wrong fast")
	}
	if avgGasPrice.Fastest != 11 {
		t.Error("wrong fastest")
	}
}

func TestBatchRequests(t *testing.T) {
	providers := []common.Provider{
		tester.NewProvider(true, common.GasPrice{}),
		tester.NewProvider(false, common.GasPrice{}),
	}

	batchResults := common.BatchRequests(providers)

	if len(batchResults) != 2 {
		t.Error("2 results expected")
		return
	}

	var countOk, countErr uint8
	for _, br := range batchResults {
		if br.Err != nil {
			countErr++
			continue
		}

		countOk++
	}

	if countOk != 1 || countErr != 1 {
		t.Error("expected 1 ok and 1 err")
	}
}
