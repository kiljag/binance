package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

type KlineService struct {
	c         *Client
	symbol    string
	interval  string
	limit     int64
	startTime int64
	endTime   int64
}

func (c *Client) NewKlinesService(symbol, interval string, limit, startTime, endTime int64) *KlineService {
	return &KlineService{
		c:         c,
		symbol:    symbol,
		interval:  interval,
		limit:     limit,
		startTime: startTime,
		endTime:   endTime,
	}
}

// get klines from binance rest api
func (s *KlineService) GetKlines() []*Kline {

	klines := make([]*Kline, 0)
	if s.symbol == "" || s.interval == "" {
		log.Println("error in fetching klines, symbol and interval can not be empty")
		return klines
	}

	req := request{
		method:   http.MethodGet,
		endpoint: endPointKlines,
	}
	req.setParam(key_SYMBOL, s.symbol)
	req.setParam(key_INTERVAL, s.interval)
	if s.limit > 0 {
		req.setParam(key_LIMIT, s.limit)
	}
	if s.startTime > 0 && s.endTime > 0 {
		req.setParam(key_STARTTIME, s.startTime)
		req.setParam(key_ENDTIME, s.endTime)
	}

	data, err := s.c.callAPI(&req)
	if err != nil {
		return klines
	}
	// log.Println(string(data))

	var klist [][]interface{}
	err = json.Unmarshal(data, &klist)
	if err != nil {
		log.Println("error in parsing klines rest api : ", err, string(data))
		return klines
	}
	// log.Println(klist)

	for _, k := range klist {
		kline := &Kline{
			Symbol:              s.symbol,
			Interval:            s.interval,
			EventTime:           int64(k[6].(float64)),
			OpenTime:            int64(k[0].(float64)),
			OpenPrice:           parseFloat(k[1].(string)),
			HighPrice:           parseFloat(k[2].(string)),
			LowPrice:            parseFloat(k[3].(string)),
			ClosePrice:          parseFloat(k[4].(string)),
			BaseVolume:          parseFloat(k[5].(string)),
			CloseTime:           int64(k[6].(float64)),
			QuoteVolume:         parseFloat(k[7].(string)),
			TradeCount:          int64(k[8].(float64)),
			TakerBuyBaseVolume:  parseFloat(k[9].(string)),
			TakerBuyQuoteVolume: parseFloat(k[10].(string)),
		}
		// log.Println(kline)
		klines = append(klines, kline)
	}

	return klines
}
