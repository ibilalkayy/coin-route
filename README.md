# Coin-Route: Optimizing the Crypto Exchange Rates 📈

## Overview

Coin-Route is a dynamic JSON API designed to find the best cryptocurrency exchange to purchase Bitcoin (BTC) with the lowest USD or USDT expenditure. This updated README provides a comprehensive guide on the project's structure, functionalities, and integration of both Coinbase and Bitfinex exchanges.

## Project Structure

```
./
|-- LICENSE
|-- README.md
|-- backend/
|   |-- bitfinex.go
|   |-- btc_amount.go
|   |-- coinbase.go
|   |-- exchange.go
|   |-- usd_amount.go
|-- go.mod
|-- main.go
|-- structs/
    |-- structs.go
```

### Components

- `main.go`: Entry point of the application.
- `backend/`: Contains backend logic for API endpoints.
  - `bitfinex.go`: Handles API calls to Bitfinex.
  - `btc_amount.go`: Extracts BTC amount from the request.
  - `coinbase.go`: Handles API calls to Coinbase.
  - `exchange.go`: Defines API routing and responses.
  - `usd_amount.go`: Calculates the lowest USD amount for a given BTC amount.
- `structs/`: Data structures used in the application.
  - `structs.go`: Defines `Response` and `APIResponse` structs.

## API Endpoints

### Exchange Routing

- **URL**: `/exchange-routing`
- **Method**: `GET`
- **Parameters**: 
  - `amount`: Amount of Bitcoin to buy (in BTC).

#### Example

```bash
curl http://localhost:4000/exchange-routing?amount=1
```

#### Response

```json
{
  "btcAmount": 1,
  "usdAmount": 10000,
  "exchange": "coinbase"
}
```

## Detailed Working

### 1. BTC Amount Extraction

- **Function**: `TakeBTCAmount(w http.ResponseWriter, r *http.Request) float64`
  
- **Workflow**:
  - Extracts the BTC amount from the query parameter `amount`.
  - Validates the amount format and returns it as a float64 value.
  - Responds with an HTTP error if the amount is missing or invalid.

### 2. API Calls to Coinbase and Bitfinex

#### Coinbase API 📖
- Documentation: [Coinbase API Reference](https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproductbook)

#### Bitfinex API 📖
- Documentation: [Bitfinex API Reference](https://docs.bitfinex.com/reference/rest-public-ticker)

#### Bitfinex API Calls (`bitfinex.go`)
- **Function**: `getBTCUSD() (float64, error)`
  
- **Workflow**:
  - Fetches the BTC/USD rate from Bitfinex's ticker API.
  - Returns the last traded price as a float64 value.

### 3. Calculate Lowest USD Amounts

#### Coinbase (`usd_amount.go`)
- **Function**: `GiveUSDAmount(w http.ResponseWriter, btcAmount float64) float64`

- **Workflow**:
  - Fetches the order book from Coinbase using `CoinbaseAPICall()`.
  - Finds the lowest USD amount for the given BTC amount.

#### Bitfinex (`bitfinex.go`)
- **Function**: `BitfinexAPICall(btcAmount float64) float64`

- **Workflow**:
  - Calculates the USD value for the given BTC amount using the Bitfinex API.

### 4. Exchange Rate Comparison

- **Function**: `TwoAmountsComparison(w http.ResponseWriter, btcAmount float64) (float64, string)`

- **Workflow**:
  - Compares the USD amounts from Coinbase and Bitfinex.
  - Determines the exchange with the lowest USD amount.

### 5. Exchange Routing (`exchange.go`)

- **Function**: `ExchangeRouting(w http.ResponseWriter, r *http.Request)`

- **Workflow**:
  - Orchestrates the entire process by calling BTC extraction, USD calculations, and exchange rate comparison.
  - Constructs the API response with `BTCAmount`, `USDAmount`, and `exchange`.
  - Sends the JSON response with HTTP status code 200.

## Setup and Run

1. Clone the repository:

```bash
git clone https://github.com/ibilalkayy/coin-route
```

2. Run the application:

```bash
go run main.go
```

The server will start listening on port `4000`.

## Dependencies

- Go standard library
- `github.com/ibilalkayy/coin-route/structs`: Custom structs for API response

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

For further details or queries, refer to the API documentation links or the source code. Enjoy optimizing exchange rates with Coin-Route! 🚀

## Dependencies

- Go standard library
- `github.com/ibilalkayy/coin-route/structs`: Custom structs for API response

## License

This project is licensed under the Apache-2.0 License. See the [LICENSE](LICENSE) file for details.