package backend

import (
	"net/http"
	"strconv"

	"github.com/safepay/safepay-test/structs"
)

func GiveUSDAmount(w http.ResponseWriter, btcAmount float64) float64 {
	apiResponse := structs.APIResponse{
		Asks: [][]string{
			{"47781.36", "0.01077963", "50c2915c-670b-4c94-8c16-fc676dcbf454"},
			{"47520", "2.15537023", "738b9f96-f6b4-4095-9bb4-1d60eafe1371"},
			{"47729.09", "0.000016", "ce1bd354-0ee0-438c-9928-3962aa638ce4"},
		},
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
