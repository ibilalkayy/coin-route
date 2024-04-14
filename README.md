# Coin-Route: Cryptocurrency Exchange Price Optimizer

## Overview

Coin-Route is a lightweight JSON API designed to help users find the best cryptocurrency exchange to buy a specified amount of Bitcoin (BTC) with minimal USD or USDT expenditure. This README provides a comprehensive guide to understanding, setting up, and extending Coin-Route.

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

- `main.go`: Entry point of the application.
- `backend/`: Backend logic for API endpoints.
  - `btc_amount.go`: Extracts BTC amount from the request.
  - `coinbase.go`: Handles API calls to Coinbase.
  - `exchange.go`: Defines API routing and responses.
  - `usd_amount.go`: Calculates lowest USD amount for a given BTC amount.
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

### 2. API Call to Coinbase

- **Function**: `CoinbaseAPICall() (*structs.APIResponse, error)`

- **Workflow**:
  - Constructs an HTTP GET request to Coinbase's API endpoint.
  - Sends the request and receives the order book response.
  - Parses the JSON response into an `APIResponse` struct.

### 3. Calculate Lowest USD Amount

- **Function**: `GiveUSDAmount(w http.ResponseWriter, btcAmount float64) float64`

- **Workflow**:
  - Calls `CoinbaseAPICall` to fetch the order book.
  - Iterates over the `Asks` in the order book to find the lowest USD amount for the given BTC amount.
  - Returns the lowest USD amount.
  - Handles potential errors such as invalid data or no suitable ask found.

### 4. Exchange Routing

- **Function**: `ExchangeRouting(w http.ResponseWriter, r *http.Request)`

- **Workflow**:
  - Calls `TakeBTCAmount` to get the BTC amount.
  - Calls `GiveUSDAmount` to calculate the lowest USD amount.
  - Constructs the API response with `BTCAmount`, `USDAmount`, and `exchange`.
  - Sends the JSON response with HTTP status code 200.

## Adding Binance Support

To support Binance, additional steps would be required:

1. Implement an API call function to fetch the order book from Binance.
2. Modify `GiveUSDAmount` to compare results from both Coinbase and Binance.
3. Update `ExchangeRouting` to determine the best exchange for the user.

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
- `github.com/safepay/safepay-test/structs`: Custom structs for API response

## License

This project is licensed under the Apache-2.0 License. See the [LICENSE](LICENSE) file for details.