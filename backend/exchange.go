package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ibilalkayy/coin-route/structs"
)

func TwoAmountsComparison(w http.ResponseWriter, btcAmount float64) (float64, string) {
	coinbaseUsdAmount := GiveUSDAmount(w, btcAmount)
	bitfinexUsdAmount := BitfinexAPICall(btcAmount)

	var min float64
	var value string

	if coinbaseUsdAmount < bitfinexUsdAmount {
		min = coinbaseUsdAmount
		value = "coinbase"
	} else {
		min = bitfinexUsdAmount
		value = "bitfinix"
	}

	fmt.Printf("Coinbase: $%.2f\n", coinbaseUsdAmount)
	fmt.Printf("Bitfinix: $%.2f\n", bitfinexUsdAmount)
	return min, value
}

func ExchangeRouting(w http.ResponseWriter, r *http.Request) {
	btcAmount := TakeBTCAmount(w, r)
	minAmount, exchange := TwoAmountsComparison(w, btcAmount)
	response := structs.Response{
		BTCAmount: btcAmount,
		USDAmount: minAmount,
		Exchange:  exchange,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
