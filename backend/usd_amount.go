package backend

import (
	"errors"
	"strconv"
)

func GetUSDFromBTC(btcAmount float64) (float64, error) {
	apiResponsePtr, err := CoinbaseAPICall()
	if err != nil {
		return 0, err
	}

	// Check if the API response has at least one ask entry
	if len(apiResponsePtr.Asks) < 1 {
		return 0, errors.New("unexpected API response format")
	}

	// Extract the USD and BTC amounts from the first ask entry
	usdStr := apiResponsePtr.Asks[0][0]
	btcStr := apiResponsePtr.Asks[0][1]

	// Convert the USD and BTC amounts from string to float64
	usdAmount, err := strconv.ParseFloat(usdStr, 64)
	if err != nil {
		return 0, errors.New("failed to parse USD amount")
	}

	btcAPIAmount, err := strconv.ParseFloat(btcStr, 64)
	if err != nil {
		return 0, errors.New("failed to parse BTC amount")
	}

	// Calculate USD amount based on the given BTC amount
	if btcAPIAmount == 0 {
		return 0, errors.New("BTC amount should not be zero")
	}

	usdAmountFromBTC := usdAmount / btcAPIAmount * btcAmount
	return usdAmountFromBTC, nil
}
