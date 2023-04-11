package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *Client) GetExchangeInfo() *ExchangeInfo {

	req := request{
		method:   http.MethodGet,
		endpoint: endPointExchangeInfo,
		query:    nil,
	}
	data, err := c.callAPI(&req)
	if err != nil {
		return &ExchangeInfo{}
	}
	var exchangeInfo ExchangeInfo
	err = json.Unmarshal(data, &exchangeInfo)
	if err != nil {
		log.Println("error in parsing exchange info", err, string(data))
		return &ExchangeInfo{}
	}

	return &exchangeInfo
}

type ExchangeInfo struct {
	ExchangeFilter interface{}     `json:"exchangeFilters"`
	RateLimits     []InfoRateLimit `json:"rateLimits"`
	ServerTime     int64           `json:"serverTime"`
	Assets         []InfoAsset     `json:"assets"`
	Symbols        []InfoSymbol    `json:"symbols"`
}

type InfoRateLimit struct {
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type InfoAsset struct {
	Asset           string `json:"asset"`
	MarginAvailable bool   `json:"marginAvailable"`
}

type InfoSymbol struct {
	Symbol                   string `json:"symbol"`
	Pair                     string `json:"pair"`
	ContractType             string `json:"contractType"`
	DeliveryDate             int64  `json:"deliveryDate"`
	OnboardDate              int64  `json:"onboardDate"`
	Status                   string `json:"status"`
	MaintenanceMarginPercent string `json:"maintMarginPercent"`
	RequiredMarginPercent    string `json:"requiredMarginPercent"`
	BaseAsset                string `json:"baseAsset"`
	QuoteAsset               string `json:"quoteAsset"`
	MarginAsset              string `json:"marginAsset"`
	PricePrecision           int64  `json:"pricePrecision"`
	QuantityPrecision        int64  `json:"quantityPrecision"`
	BaseAssetPrecision       int64  `json:"baseAssetPrecision"`
	QuotePrecision           int64  `json:"quotePrecision"`
}
