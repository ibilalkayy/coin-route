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

	var lowestUSDAmount float64
	lowestUSDAmountFound := false

	for _, ask := range apiResponse.Asks {
		askBTC, err := strconv.ParseFloat(ask[1], 64)
		if err != nil {
			continue
		}

		askUSD, err := strconv.ParseFloat(ask[0], 64)
		if err != nil {
			http.Error(w, "Invalid USD amount format", http.StatusInternalServerError)
			return 0
		}

		if askBTC >= btcAmount {
			if !lowestUSDAmountFound || askUSD < lowestUSDAmount {
				lowestUSDAmount = askUSD
				lowestUSDAmountFound = true
			}
		}
	}

	if !lowestUSDAmountFound {
		http.Error(w, "No suitable ask found for the specified BTC amount", http.StatusInternalServerError)
		return 0
	}

	return lowestUSDAmount
}
