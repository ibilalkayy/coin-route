package backend

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ibilalkayy/coin-route/structs"
)

func CoinbaseAPICall() (*structs.APIResponse, error) {
	url := "https://api-public.sandbox.exchange.coinbase.com/products/BTC-USD/book?level=3"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse structs.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
