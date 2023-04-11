package binance

import (
	"encoding/json"
	"fmt"
	"log"
)

type TickerStream struct {
	c   *Client
	out chan *PriceTicker
	wss *WebSocketStream
}

type jsonPriceTickerEvent struct {
	EventType          string `json:"e"`
	EventTime          int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	LastPrice          string `json:"c"`
	LastQuantity       string `json:"Q"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstTradeId       int64  `json:"F"`
	LastTradeId        int64  `json:"L"`
	TradeCount         int64  `json:"n"`
}

func (c *Client) NewTickerStream() *TickerStream {
	url := fmt.Sprintf("%s/!ticker@arr", baseWsMainURL)
	return &TickerStream{
		c:   c,
		out: make(chan *PriceTicker, 100),
		wss: &WebSocketStream{
			url:     url,
			timeout: 2 * 60 * 60 * 1000,
		},
	}
}

func (s *TickerStream) Start() <-chan *PriceTicker {
	s.wss.isActive = true
	go s.startStream()
	return s.out
}

func (s *TickerStream) Stop() {
	s.wss.isActive = false
}

func (s *TickerStream) startStream() {

	defer close(s.out)
	messageCount := 0

	for {
		if !s.wss.isActive {
			break
		}
		msg, err := s.wss.getNextMessage()
		if err != nil {
			break
		}

		var eventList []jsonPriceTickerEvent
		err = json.Unmarshal(msg, &eventList)
		if err != nil {
			log.Println("error in parsing ws ticker : ", err)
		}

		for _, event := range eventList {
			ticker := PriceTicker{
				Symbol:             event.Symbol,
				PriceChange:        parseFloat(event.PriceChange),
				PriceChangePercent: parseFloat(event.PriceChangePercent),
				WeightedAvgPrice:   parseFloat(event.WeightedAvgPrice),
				LastPrice:          parseFloat(event.LastPrice),
				LastQuantity:       parseFloat(event.LastQuantity),
				OpenPrice:          parseFloat(event.OpenPrice),
				HighPrice:          parseFloat(event.HighPrice),
				LowPrice:           parseFloat(event.LowPrice),
				BaseVolume:         parseFloat(event.BaseVolume),
				QuoteVolume:        parseFloat(event.QuoteVolume),
				OpenTime:           event.OpenTime,
				CloseTime:          event.CloseTime,
				FirstTradeId:       event.FirstTradeId,
				LastTradeId:        event.LastTradeId,
				TradeCount:         event.TradeCount,
			}
			select {
			case s.out <- &ticker:
				messageCount += 1
			default:
				log.Println("error in ticker stream : out channel is full")
			}
		}
	}
}
