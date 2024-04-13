package structs

type Response struct {
	BTCAmount float64 `json:"btcAmount"`
	USDAmount float64 `json:"usdAmount"`
	Exchange  string  `json:"exchange"`
}

type APIResponse struct {
	Asks [][]string `json:"asks"`
}
