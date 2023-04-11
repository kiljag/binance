package binance

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

/* stream to get updates on orders, positions etc.*/

type AccountStream struct {
	c         *Client
	listenKey string
	out       chan interface{}

	isActive   bool
	wsConn     *websocket.Conn
	wsOpenTime int64
	timeout    int64
}

func (c *Client) NewAccountStream() *AccountStream {
	return &AccountStream{
		c:       c,
		out:     make(chan interface{}),
		timeout: 30 * 60 * 1000, // refresh ws every 30 minutes
	}
}

func (s *AccountStream) Start() <-chan interface{} {
	s.isActive = true
	go s.startStream()
	return s.out
}

func (s *AccountStream) Stop() {
	s.isActive = false
}

func (s *AccountStream) getNextMessage() ([]byte, error) {
	if !s.isActive {
		return nil, fmt.Errorf("account stream is inactive")
	}

	if s.wsOpenTime == 0 || (CurrentTimestamp()-s.wsOpenTime) > s.timeout {
		if s.wsConn != nil {
			s.wsConn.Close()
		}
		listenKey := s.getListenKey()
		if listenKey == "" {
			log.Println("error invalid listen key, closing account stream")
			return nil, fmt.Errorf("wstream error")
		}
		s.listenKey = listenKey
		url := fmt.Sprintf("%s/%s", baseWsMainURL, s.listenKey)
		log.Println("opening new wstream: " + url)
		c, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Println("error in creating account wstream : ", err)
			return nil, err
		}
		s.wsConn = c
		s.wsOpenTime = CurrentTimestamp()
	}
	_, msg, err := s.wsConn.ReadMessage()
	if err != nil {
		log.Println("error in reading account wstream : ", err)
		return nil, err
	}
	return msg, err
}

func (s *AccountStream) startStream() {
	defer close(s.out)
	messageCount := 0
	for {
		msg, err := s.getNextMessage()
		if s.c.debug {
			log.Println("==> event : " + string(msg))
		}
		if err != nil {
			break
		}
		event := s.parseResponse(msg)
		if event == nil {
			continue
		}
		s.out <- event
		messageCount += 1
	}
}

func (s *AccountStream) parseResponse(data []byte) interface{} {
	var accountEvent interface{}
	var event map[string]interface{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		log.Println("error in parsing account ws response : ", err, string(data))
		return nil
	}

	eventType := event["e"].(string)
	if s.c.debug {
		log.Println("==> eventType : ", eventType)
	}

	if eventType == EVENT_MARGIN_CALL {
		accountEvent = s.parseMarginCallEvent(data)
	} else if eventType == Event_ACCOUNT_UPDATE {
		accountEvent = s.parseAccountUpdateEvent(data)
	} else if eventType == Event_ORDER_TRADE_UPDATE {
		accountEvent = s.parseOrderTradeUpdateEvent(data)
	}

	return accountEvent
}

// margin call event
type jsonMCPosition struct {
	Symbol            string `json:"s"`
	PositionSide      string `json:"ps"`
	PositionAmount    string `json:"pa"`
	MarginType        string `json:"mt"`
	IsolatedWallet    string `json:"iw"`
	MarketPrice       string `json:"mp"`
	UnrealizedPnL     string `json:"up"`
	MaintenanceMargin string `json:"mm"`
}

type jsonMarginCallEvent struct {
	Event              string           `json:"e"`
	EventTime          int64            `json:"E"`
	CrossWalletBalance string           `json:"cw"`
	Positions          []jsonMCPosition `json:"p"`
}

func (s *AccountStream) parseMarginCallEvent(data []byte) *MarginCallEvent {
	var event jsonMarginCallEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		log.Println("error in parsing margin call event : ", err, string(data))
		return nil
	}
	positions := make([]MarginPosition, 0)
	for _, p := range event.Positions {
		position := MarginPosition{
			Symbol:            p.Symbol,
			PositionSide:      p.PositionSide,
			PositionAmount:    parseFloat(p.PositionAmount),
			MarginType:        p.MarginType,
			IsolatedWallet:    parseFloat(p.IsolatedWallet),
			MarketPrice:       parseFloat(p.MarketPrice),
			UnrealizedPnL:     parseFloat(p.UnrealizedPnL),
			MaintenanceMargin: parseFloat(p.MaintenanceMargin),
		}
		positions = append(positions, position)
	}
	return &MarginCallEvent{
		Event:              event.Event,
		EventTime:          event.EventTime,
		CrossWalletBalance: parseFloat(event.CrossWalletBalance),
		Positions:          positions,
	}
}

// account update event
type jsonACBalance struct {
	Asset              string `json:"a"`
	WalletBalance      string `json:"wb"`
	CrossWalletBalance string `json:"cw"`
	BalanceChange      string `json:"bc"`
}

type jsonACPosition struct {
	Symbol         string `json:"s"`
	PositionAmount string `json:"pa"`
	EntryPrice     string `json:"ep"`
	Accumulated    string `json:"cr"`
	UnrealizedPnL  string `json:"up"`
	MarginType     string `json:"mt"`
	IsolatedWallet string `json:"iw"`
	PositionSide   string `json:"ps"`
}

type jsonACData struct {
	UpdateType string           `json:"m"`
	Balances   []jsonACBalance  `json:"B"`
	Positions  []jsonACPosition `json:"P"`
}

type jsonAccountUpdateEvent struct {
	Event      string     `json:"e"`
	EventTime  int64      `json:"E"`
	Transation int64      `json:"T"`
	UpdateData jsonACData `json:"a"`
}

func (s *AccountStream) parseAccountUpdateEvent(data []byte) *AccountUpdateEvent {
	event := new(jsonAccountUpdateEvent)
	err := json.Unmarshal(data, event)
	if err != nil {
		log.Println("error in parsing account update event : ", err, string(data))
		return nil
	}
	balances := make([]AccountUpdateBalance, 0)
	for _, b := range event.UpdateData.Balances {
		balance := AccountUpdateBalance{
			Asset:              b.Asset,
			WalletBalance:      parseFloat(b.WalletBalance),
			CrossWalletBalance: parseFloat(b.CrossWalletBalance),
			BalanceChange:      parseFloat(b.CrossWalletBalance),
		}
		balances = append(balances, balance)
	}

	positions := make([]AccountUpdatePosition, 0)
	for _, p := range event.UpdateData.Positions {
		position := AccountUpdatePosition{
			Symbol:         p.Symbol,
			PositionAmount: parseFloat(p.PositionAmount),
			EntryPrice:     parseFloat(p.EntryPrice),
			Accumulated:    parseFloat(p.Accumulated),
			UnrealizedPnL:  parseFloat(p.UnrealizedPnL),
			MarginType:     p.MarginType,
			IsolatedWallet: parseFloat(p.IsolatedWallet),
			PositionSide:   p.PositionSide,
		}
		positions = append(positions, position)
	}

	return &AccountUpdateEvent{
		Event:           event.Event,
		EventTime:       event.EventTime,
		TransactionTime: event.Transation,
		UpdateData: AccountUpdateData{
			UpdateType: event.UpdateData.UpdateType,
			Balances:   balances,
			Positions:  positions,
		},
	}
}

// order trade update events
type jsonOrderTradeData struct {
	Symbol               string `json:"s"`
	ClientOrderId        string `json:"c"`
	OrderSide            string `json:"S"`
	OrderType            string `json:"o"`
	TimeInForce          string `json:"f"`
	Quantity             string `json:"q"`
	Price                string `json:"p"`
	AveragePrice         string `json:"ap"`
	StopPrice            string `json:"sp"`
	ExectutionType       string `json:"x"`
	OrderStatus          string `json:"X"`
	OrderId              int64  `json:"i"`
	LastFilledQuantity   string `json:"l"`
	AccumulatedQuantity  string `json:"z"`
	LastFilledPrice      string `json:"L"`
	CommissionAsset      string `json:"N"`
	Commission           string `json:"n"`
	TradeTime            int64  `json:"T"`
	TradeId              int64  `json:"t"`
	BidsNotional         string `json:"b"`
	AskNotional          string `json:"a"`
	IsMakerSide          bool   `json:"m"`
	IsReduceOnly         bool   `json:"R"`
	StopPriceWorkingType string `json:"wt"`
	OringalOrderType     string `json:"ot"`
	PositionSide         string `json:"ps"`
	IsCloseAll           bool   `json:"cp"`
	ActivationPrice      string `json:"AP"`
	CallbackRate         string `json:"cr"`
	RealizedProfit       string `json:"rp"`
}

type jsonOrderTradeUpdateEvent struct {
	Event          string             `json:"e"`
	EventTime      int64              `json:"E"`
	TransationTime int64              `json:"T"`
	OrderData      jsonOrderTradeData `json:"o"`
}

func (s *AccountStream) parseOrderTradeUpdateEvent(data []byte) *OrderTradeUpdateEvent {
	event := new(jsonOrderTradeUpdateEvent)
	err := json.Unmarshal(data, event)
	if err != nil {
		log.Println("error in parsing trade update event : ", err, string(data))
		return nil
	}
	order := event.OrderData

	return &OrderTradeUpdateEvent{
		Event:          event.Event,
		EventTime:      event.EventTime,
		TransationTime: event.TransationTime,
		OrderData: OrderTradeData{
			Symbol:               order.Symbol,
			ClientOrderId:        order.ClientOrderId,
			OrderSide:            order.OrderSide,
			OrderType:            order.OrderType,
			TimeInForce:          order.TimeInForce,
			Quantity:             parseFloat(order.Quantity),
			Price:                parseFloat(order.Price),
			AveragePrice:         parseFloat(order.AveragePrice),
			StopPrice:            parseFloat(order.StopPrice),
			ExectutionType:       order.ExectutionType,
			OrderStatus:          order.OrderStatus,
			OrderId:              order.OrderId,
			LastFilledQuantity:   parseFloat(order.LastFilledQuantity),
			LastFilledPrice:      parseFloat(order.LastFilledPrice),
			CommissionAsset:      order.CommissionAsset,
			Commission:           parseFloat(order.Commission),
			TradeTime:            order.TradeTime,
			TradeId:              order.TradeId,
			BidsNotional:         parseFloat(order.BidsNotional),
			AskNotional:          parseFloat(order.AskNotional),
			IsMakerSide:          order.IsMakerSide,
			IsReduceOnly:         order.IsReduceOnly,
			StopPriceWorkingType: order.StopPriceWorkingType,
			OringalOrderType:     order.OringalOrderType,
			PositionSide:         order.PositionSide,
			IsCloseAll:           order.IsCloseAll,
			ActivationPrice:      parseFloat(order.ActivationPrice),
			CallbackRate:         parseFloat(order.CallbackRate),
			RealizedProfit:       parseFloat(order.RealizedProfit),
		},
	}
}
