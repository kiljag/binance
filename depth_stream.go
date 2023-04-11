package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DepthStream struct {
	c      *Client
	symbol string
	level  int // number of book entries
	out    chan *OrderBookEvent
	wss    *WebSocketStream
}

func (c *Client) NewDepthStream(symbol string, level int) *DepthStream {
	endpoint := fmt.Sprintf("%s@depth%s", strings.ToLower(symbol), strconv.Itoa(level))
	url := fmt.Sprintf("%s/%s", baseWsMainURL, endpoint)

	return &DepthStream{
		c:      c,
		symbol: symbol,
		level:  level,
		out:    make(chan *OrderBookEvent),
		wss: &WebSocketStream{
			url:     url,
			timeout: 2 * 60 * 60 * 1000,
		},
	}
}

func (s *DepthStream) Start() <-chan *OrderBookEvent {
	s.wss.isActive = true
	go s.startStream()
	return s.out
}

func (s *DepthStream) Stop() {
	s.wss.isActive = false
}

func (s *DepthStream) startStream() {
	defer close(s.out)
	messageCount := 0
	for {
		msg, err := s.wss.getNextMessage()
		if err != nil {
			break
		}
		event := s.parseResponse(msg)
		if event == nil {
			log.Println("error order book event is nil")
			continue
		}

		s.out <- event
		messageCount += 1
	}
	log.Printf("sent %d depth events for %s", messageCount, s.symbol)
}

type jsonDepthEvent struct {
	EventType      string          `json:"e"`
	EventTime      int64           `json:"E"`
	TransationTime int64           `json:"T"`
	Symbol         string          `json:"s"`
	Bids           [][]interface{} `json:"b"`
	Asks           [][]interface{} `json:"a"`
}

func (s *DepthStream) parseResponse(data []byte) *OrderBookEvent {
	var event jsonDepthEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		log.Println("error in parsing depth event : ", err, string(data))
		return nil
	}

	bids := make([]OrderBookEntry, 0)
	for _, b := range event.Bids {
		BidEntry := OrderBookEntry{
			Price:    parseFloat(b[0].(string)),
			Quantity: parseFloat(b[1].(string)),
		}
		bids = append(bids, BidEntry)
	}

	asks := make([]OrderBookEntry, 0)
	for _, a := range event.Asks {
		AskEntry := OrderBookEntry{
			Price:    parseFloat(a[0].(string)),
			Quantity: parseFloat(a[1].(string)),
		}
		asks = append(asks, AskEntry)
	}

	return &OrderBookEvent{
		EventType:       event.EventType,
		EventTime:       event.EventTime,
		TransactionTime: event.TransationTime,
		Symbol:          event.Symbol,
		Bids:            bids,
		Asks:            asks,
	}
}
