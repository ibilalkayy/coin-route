package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/safepay/safepay-test/backend"
)

func Execute() error {
	http.HandleFunc("/exchange-routing", backend.ExchangeRouting)
	fmt.Println("Listening the server at port: 4000")
	return http.ListenAndServe(":4000", nil)
}

func main() {
	err := Execute()
	if err != nil {
		log.Fatal(err)
	}
}
