package backend

import (
	"net/http"
	"strconv"
)

func TakeBTCAmount(w http.ResponseWriter, r *http.Request) float64 {
	amountStr := r.URL.Query().Get("amount")
	if amountStr == "" {
		http.Error(w, "Amount query parameter is missing", http.StatusBadRequest)
		return 0
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return 0
	}
	return amount
}
