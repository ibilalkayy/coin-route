package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getBTCUSD() (float64, error) {
	resp, err := http.Get("https://api-pub.bitfinex.com/v2/ticker/tBTCUSD")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Parsing the JSON response
	var data []interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	// The current BTC/USD rate is the last traded price in the response
	rate, ok := data[6].(float64)
	if !ok {
		return 0, fmt.Errorf("failed to parse rate")
	}

	return rate, nil
}

func calculateBTCValueInUSD(btcAmount float64) (float64, error) {
	rate, err := getBTCUSD()
	if err != nil {
		return 0, err
	}

	usdValue := btcAmount * rate
	return usdValue, nil
}

func BitfinexAPICall(btcAmount float64) float64 {
	usdValue, err := calculateBTCValueInUSD(btcAmount)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	// fmt.Printf("%.2f BTC is equivalent to $%.2f\n", btcAmount, usdValue)
	return usdValue
}
