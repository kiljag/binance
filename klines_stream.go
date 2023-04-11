package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type KlineStream struct {
	c        *Client
	symbol   string
	interval string
	out      chan *Kline
	wss      *WebSocketStream
	dropProb float32
}

func (c *Client) NewKlineStream(symbol string, interval string, dropProb float32) *KlineStream {
	if symbol == "" || interval == "" || dropProb < 0 || dropProb > 1 {
		log.Println("error in kline stream, empty symbol or interval")
	}
	// endpoint := fmt.Sprintf("%s_perpetual@continuousKline_%s", strings.ToLower(symbol), interval)
	endpoint := fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval)
	url := fmt.Sprintf("%s/%s", baseWsMainURL, endpoint)

	return &KlineStream{
		c:        c,
		symbol:   symbol,
		interval: interval,
		out:      make(chan *Kline),
		wss: &WebSocketStream{
			url:     url,
			timeout: 2 * 60 * 60 * 1000,
		},
		dropProb: dropProb,
	}
}

func (s *KlineStream) Start() <-chan *Kline {
	s.wss.isActive = true
	go s.startStream()
	return s.out
}

// func (s *KlineStream) setDropProbablity(dropProb float32) {
// 	s.dropProb = dropProb
// }

func (s *KlineStream) Stop() {
	s.wss.isActive = false
}

func (s *KlineStream) startStream() {
	defer close(s.out)
	messageCount := 0

	for {
		msg, err := s.wss.getNextMessage()
		if err != nil {
			break
		}

		kline := s.parseResponse(msg)
		if rand.Float32() < s.dropProb && !kline.IsFinal {
			continue
		}

		if kline == nil {
			log.Println("error, kline event is nil", s.symbol, kline)
			continue
		}
		messageCount += 1
		s.out <- kline
	}
}

type jsonWsKline struct {
	StartTime           int64  `json:"t"`
	EndTime             int64  `json:"T"`
	Symbol              string `json:"s"`
	Interval            string `json:"i"`
	FirstTradeID        int64  `json:"f"`
	LastTradeID         int64  `json:"L"`
	Open                string `json:"o"`
	Close               string `json:"c"`
	High                string `json:"h"`
	Low                 string `json:"l"`
	BaseVolume          string `json:"v"`
	TradeCount          int64  `json:"n"`
	IsFinal             bool   `json:"x"`
	QuoteAssetVolume    string `json:"q"`
	TakerBuyBaseVolume  string `json:"V"`
	TakerBuyQuoteVolume string `json:"Q"`
}

type jsonWsKlineEvent struct {
	Event  string      `json:"e"`
	Time   int64       `json:"E"`
	Symbol string      `json:"s"`
	Kline  jsonWsKline `json:"k"`
}

func (s *KlineStream) parseResponse(data []byte) *Kline {
	var event jsonWsKlineEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		log.Println("error in parsing kline : ", err, string(data))
		return nil
	}
	k := event.Kline

	return &Kline{
		Symbol:              k.Symbol,
		Interval:            k.Interval,
		EventTime:           event.Time,
		OpenTime:            k.StartTime,
		CloseTime:           k.EndTime,
		OpenPrice:           parseFloat(k.Open),
		HighPrice:           parseFloat(k.High),
		LowPrice:            parseFloat(k.Low),
		ClosePrice:          parseFloat(k.Close),
		BaseVolume:          parseFloat(k.BaseVolume),
		QuoteVolume:         parseFloat(k.QuoteAssetVolume),
		TradeCount:          event.Kline.TradeCount,
		TakerBuyBaseVolume:  parseFloat(k.TakerBuyBaseVolume),
		TakerBuyQuoteVolume: parseFloat(k.TakerBuyQuoteVolume),
		IsFinal:             k.IsFinal,
	}
}
