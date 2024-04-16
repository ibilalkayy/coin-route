package backend

import (
	"net/http"
	"strconv"
)

func GiveUSDAmount(w http.ResponseWriter, btcAmount float64) float64 {
	apiResponse, err := CoinbaseAPICall()
	if err != nil {
		http.Error(w, "Failed to fetch data from Coinbase API", http.StatusInternalServerError)
		return 0
	}

	var usdAmount float64
	for _, ask := range apiResponse.Asks {
		// Convert the BTC amount from string to float64
		askBTC, err := strconv.ParseFloat(ask[1], 64)
		if err != nil {
			continue
		}

		if askBTC == btcAmount {
			// Convert the USD amount from string to float64
			usdAmount, err = strconv.ParseFloat(ask[0], 64)
			if err != nil {
				http.Error(w, "Invalid USD amount format", http.StatusInternalServerError)
				return 0
			}
			break
		}
	}
	return usdAmount
}
