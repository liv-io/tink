package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var (
	avgPrice float64
	avgMutex sync.RWMutex
)

func getPriceData(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request (%q): %w", url, err)

	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request (%q): %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to make HTTP request (%q): %d:", url, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode JSON response (%q): %w", url, err)
	}

	return nil
}

func validatePrice(name string, price interface{}) error {
	if val, ok := price.(float64); ok {
		if val <= 0 {
			return fmt.Errorf("invalid price value in (%s) response: %f", name, val)
		}
		return nil
	}

	return fmt.Errorf("invalid type, price is not float64 in (%s) response: %v", name, price)
}

// bitfinex.com
func fetchBitfinexPrice() (float64, error) {
	type ApiResponse []float64

	name := "Bitfinex"
	url := "https://api.bitfinex.com/v2/ticker/tBTCUSD"
	var resp ApiResponse

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	price := resp[6]
	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// bitget.com
func fetchBitgetPrice() (float64, error) {
	type ApiResponse struct {
		Data []struct {
			LastPr string `json:"lastPr"`
		} `json:"data"`
	}

	name := "Bitget"
	url := "https://api.bitget.com/api/v2/spot/market/tickers?symbol=BTCUSDT"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Data) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Data[0].LastPr
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// bitrue.com
func fetchBitruePrice() (float64, error) {
	type ApiResponse struct {
		Price string `json:"price"`
	}

	name := "Bitrue"
	url := "https://openapi.bitrue.com/api/v1/ticker/price?symbol=btcusdt"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Price) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Price
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// btse.com
func fetchBTSEPrice() (float64, error) {
	type ApiResponse struct {
		LastPrice float64 `json:"lastPrice"`
	}

	name := "BTSE"
	url := "https://api.btse.com/spot/api/v3.2/price?symbol=BTC-USD"
	var resp []ApiResponse

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	price := resp[0].LastPrice

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// coinbase.com
func fetchCoinbasePrice() (float64, error) {
	type ApiResponse struct {
		Price string `json:"price"`
	}

	name := "Coinbase"
	url := "https://api.exchange.coinbase.com/products/BTC-USD/ticker"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Price) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Price
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// crypto.com
func fetchCryptoPrice() (float64, error) {
	type ApiResponse struct {
		Result struct {
			Data []struct {
				A string `json:"a"`
			} `json:"data"`
		} `json:"result"`
	}

	name := "Crypto"
	url := "https://api.crypto.com/v2/public/get-ticker?instrument_name=BTC_USDT"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Result.Data) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Result.Data[0].A
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// gateapi.io
func fetchGatePrice() (float64, error) {
	type ApiResponse struct {
		Last string `json:"last"`
	}

	name := "Gate"
	url := "https://data.gateapi.io/api2/1/ticker/sbtc_usdt"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Last) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Last
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// huobi.pro
func fetchHuobiPrice() (float64, error) {
	type ApiResponse struct {
		Tick struct {
			Data []struct {
				Price float64 `json:"price"`
			} `json:"data"`
		} `json:"tick"`
	}

	name := "Huobi"
	url := "https://api.huobi.pro/market/trade?symbol=btcusdt"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Tick.Data) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	price := resp.Tick.Data[0].Price
	err := validatePrice(name, price)
	if err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// kraken.com
func fetchKrakenPrice() (float64, error) {
	type ApiResponse struct {
		Result struct {
			XXBTZUSD struct {
				A []string `json:"a"`
			} `json:"XXBTZUSD"`
		} `json:"result"`
	}

	name := "Kraken"
	url := "https://api.kraken.com/0/public/Ticker?pair=BTCUSD"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Result.XXBTZUSD.A) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Result.XXBTZUSD.A[0]
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// kucoin.com
func fetchKuCoinPrice() (float64, error) {
	type ApiResponse struct {
		Data struct {
			Price string `json:"price"`
		} `json:"data"`
	}

	name := "KuCoin"
	url := "https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=BTC-USDC"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Data.Price) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Data.Price
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// okx.com
func fetchOKXPrice() (float64, error) {
	type ApiResponse struct {
		Data []struct {
			Last string `json:"last"`
		} `json:"data"`
	}

	name := "OKX"
	url := "https://www.okx.com/api/v5/market/ticker?instId=BTC-USDC"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Data) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Data[0].Last
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if err := validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

// xt.com
func fetchXTPrice() (float64, error) {
	type ApiResponse struct {
		Result []struct {
			C string `json:"c"`
		} `json:"result"`
	}

	name := "XT"
	url := "https://sapi.xt.com/v4/public/ticker?symbol=BTC_usdt"
	resp := ApiResponse{}

	if err := getPriceData(url, &resp); err != nil {
		return 0, fmt.Errorf("failed to make HTTP request (%s): %w", name, err)
	}

	if len(resp.Result) == 0 {
                return 0, fmt.Errorf("failed to access price field in response (%s)", name)
	}

	priceStr := resp.Result[0].C
	if priceStr == "" {
		return 0, fmt.Errorf("empty price field in response (%s)", name)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
                return 0, fmt.Errorf("failed to convert price from string to float: %w", name, err)
	}

	if validatePrice(name, price); err != nil {
		return 0, fmt.Errorf("invalid price value in response (%s): %f", name, price)
	}

	return price, nil
}

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	go func() {
		for {
			total := 0.0
			count := 0

			prices := []func() (float64, error){
				fetchBitfinexPrice,
				fetchBitgetPrice,
				fetchBitruePrice,
				fetchBTSEPrice,
				fetchCoinbasePrice,
				fetchCryptoPrice,
				fetchGatePrice,
				fetchHuobiPrice,
				fetchKrakenPrice,
				fetchKuCoinPrice,
				fetchOKXPrice,
				fetchXTPrice,
			}

			for _, fetch := range prices {
				price, err := fetch()
				if err == nil {
					total += price
					count++
				}
			}

			if count > 0 {
				avg := total / float64(count)
				avgMutex.Lock()
				avgPrice = avg
				avgMutex.Unlock()
			}

			time.Sleep(2 * time.Minute)
		}
	}()

	e.GET("/", func(c *echo.Context) error {
		avgMutex.RLock()
		defer avgMutex.RUnlock()
		return c.JSON(200, map[string]float64{"average_price": avgPrice})
	})

	e.GET("/version", func(c *echo.Context) error {
		return c.String(http.StatusOK, "0.1.0")
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
