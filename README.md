# Safepay Backend Engineer Takehome Test

## Overview

This project is a takehome test for candidates applying for a backend engineer position at Safepay. The objective is to create a JSON API that determines the best cryptocurrency exchange (Coinbase or Binance) to buy a specified amount of Bitcoin (BTC) in terms of US Dollars (USD) or Tether (USDT). This README provides an in-depth explanation of the project's structure, functionality, and implementation details.

## Project Structure

```
./
|-- LICENSE
|-- README.md
|-- backend/
|   |-- btc_amount.go
|   |-- coinbase.go
|   |-- exchange.go
|   |-- usd_amount.go
|-- database/
|-- go.mod
|-- main.go
|-- structs/
    |-- structs.go
```

### Components

- `main.go`: The entry point of the application.
- `backend/`: Contains backend logic.
  - `btc_amount.go`: Extracts the BTC amount from the API request.
  - `coinbase.go`: Makes API calls to Coinbase exchange.
  - `exchange.go`: Defines the routing and API response structure.
  - `usd_amount.go`: Calculates the lowest USD amount for a given BTC amount.
- `structs/`: Contains data structures.
  - `structs.go`: Defines `Response` and `APIResponse` structs.

## API Endpoints

### Exchange Routing

- **URL**: `/exchange-routing`
- **Method**: `GET`
- **Parameters**: 
  - `amount`: The amount of Bitcoin to buy (in BTC).

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

### 1. BTC Amount Extraction (`btc_amount.go`)

- **Function**: `TakeBTCAmount(w http.ResponseWriter, r *http.Request) float64`
  
- **Workflow**:
  - Extracts the BTC amount from the query parameter `amount`.
  - Validates the amount format and returns it as a float64 value.
  - If the amount is missing or invalid, it responds with an appropriate HTTP error.

### 2. API Call to Coinbase (`coinbase.go`)

- **Function**: `CoinbaseAPICall() (*structs.APIResponse, error)`

- **Workflow**:
  - Constructs an HTTP GET request to Coinbase's API endpoint.
  - Sends the request and receives the order book response.
  - Parses the JSON response into an `APIResponse` struct defined in `structs.go`.

### 3. Calculate Lowest USD Amount (`usd_amount.go`)

- **Function**: `GiveUSDAmount(w http.ResponseWriter, btcAmount float64) float64`

- **Workflow**:
  - Calls `CoinbaseAPICall` to get the order book.
  - Iterates over the `Asks` in the order book to find the lowest USD amount for the given BTC amount.
  - Returns the lowest USD amount.
  - Handles potential errors such as invalid data or no suitable ask found.

### 4. Exchange Routing (`exchange.go`)

- **Function**: `ExchangeRouting(w http.ResponseWriter, r *http.Request)`

- **Workflow**:
  - Calls `TakeBTCAmount` to get the BTC amount.
  - Calls `GiveUSDAmount` to get the lowest USD amount.
  - Constructs the API response with `BTCAmount`, `USDAmount`, and `exchange` (currently set to "coinbase").
  - Sends the JSON response with the HTTP status code 200.

## Adding Binance Support

To support Binance, you would need to:

1. Implement a similar API call function to fetch the order book from Binance.
2. Modify `GiveUSDAmount` to compare the results from both Coinbase and Binance to find the best USD amount.
3. Update `ExchangeRouting` to incorporate the logic for comparing and determining the best exchange.

## Setup and Run

1. Clone the repository:

```bash
git clonehttps://github.com/ibilalkayy/safepay-test
```

2. Run the application:

```bash
go run main.go
```

The server will start listening on port `4000`.

## Dependencies

- Go standard library
- `github.com/safepay/safepay-test/structs`: Custom structs for API response

## License

This project is licensed under the Apache-2.0 License. See the [LICENSE](LICENSE) file for details.

---

If you have any questions or need further clarification, feel free to reach out! 😊