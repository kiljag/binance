package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

type TickerService struct {
	c *Client
}

func (c *Client) NewTickerService() *TickerService {
	return &TickerService{
		c: c,
	}
}

func (m *TickerService) GetPriceTickers() []PriceTicker {
	req := request{
		method:   http.MethodGet,
		endpoint: endPoint24hrTicker,
		query:    nil,
	}
	data, err := m.c.callAPI(&req)
	if err != nil {
		return []PriceTicker{}
	}

	var tickerList []jsonTicker24hr
	err = json.Unmarshal(data, &tickerList)
	if err != nil {
		log.Println("error in parsing tickers response : ", err, string(data))
		return []PriceTicker{}
	}
	// log.Println("#tickers : ", len(tickers))

	tickers := make([]PriceTicker, 0)
	for _, t := range tickerList {
		ticker := PriceTicker{
			Symbol:             t.Symbol,
			PriceChange:        parseFloat(t.PriceChange),
			PriceChangePercent: parseFloat(t.PriceChangePercent),
			WeightedAvgPrice:   parseFloat(t.WeightedAvgPrice),
			LastPrice:          parseFloat(t.LastPrice),
			LastQuantity:       parseFloat(t.LastQuantity),
			OpenPrice:          parseFloat(t.OpenPrice),
			HighPrice:          parseFloat(t.HighPrice),
			LowPrice:           parseFloat(t.LowPrice),
			BaseVolume:         parseFloat(t.Volume),
			QuoteVolume:        parseFloat(t.QuoteVolume),
			OpenTime:           t.OpenTime,
			CloseTime:          t.CloseTime,
			FirstTradeId:       t.FirstTradeId,
			LastTradeId:        t.LastTradeId,
			TradeCount:         t.TradeCount,
		}
		tickers = append(tickers, ticker)
	}
	return tickers
}

type jsonTicker24hr struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	LastPrice          string `json:"lastPrice"`
	LastQuantity       string `json:"lastQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstTradeId       int64  `json:"firstId"`
	LastTradeId        int64  `json:"lastId"`
	TradeCount         int64  `json:"count"`
}
