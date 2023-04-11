package binance

// contains all the structs exposed by the binance client
type Kline struct {
	Symbol              string
	Interval            string
	EventTime           int64
	OpenTime            int64
	CloseTime           int64
	OpenPrice           float64
	HighPrice           float64
	LowPrice            float64
	ClosePrice          float64
	TradeCount          int64
	BaseVolume          float64
	QuoteVolume         float64
	TakerBuyBaseVolume  float64
	TakerBuyQuoteVolume float64
	IsFinal             bool
}

type PriceTicker struct {
	Symbol             string
	PriceChange        float64
	PriceChangePercent float64
	WeightedAvgPrice   float64
	LastPrice          float64
	LastQuantity       float64
	OpenPrice          float64
	HighPrice          float64
	LowPrice           float64
	BaseVolume         float64
	QuoteVolume        float64
	OpenTime           int64
	CloseTime          int64
	FirstTradeId       int64
	LastTradeId        int64
	TradeCount         int64
}

type OrderBookEntry struct {
	Price    float64
	Quantity float64
}

type OrderBookEvent struct {
	EventType       string
	EventTime       int64
	TransactionTime int64
	Symbol          string
	Bids            []OrderBookEntry
	Asks            []OrderBookEntry
}

type CoinBalance struct {
	AccountAlias       string
	Asset              string
	Balance            float64
	CrossWalletBalance float64
	CrossUnPnl         float64
	AvailableBalance   float64
	MaxWithdrawAmount  float64
	MarginAvailable    bool
	UpdateTime         int64
}

type LimitOrder struct {
	Symbol      string
	Side        SideType
	TimeInForce TimeInForceType
	Quantity    float64
	Price       float64
}

type MarketOrder struct {
	Symbol   string
	Side     SideType
	Quantity float64
}

type StopOrder struct {
	Symbol      string
	Side        SideType
	TimeInForce TimeInForceType
	Quantity    float64
	Price       float64
	StopPrice   float64
	ReduceOnly  bool
}

type TakeProfitOrder struct {
	Symbol     string
	Side       SideType
	Quantity   float64
	Price      float64
	StopPrice  float64
	ReduceOnly bool
}

type StopMarketOrder struct {
	Symbol     string
	Side       SideType
	StopPrice  float64
	Quantity   float64
	ReduceOnly bool
}

type TakeProfitMarketOrder struct {
	Symbol     string
	Side       SideType
	StopPrice  float64
	Quantity   float64
	ReduceOnly bool
}

type OrderResponse struct {
	ClientOrderId    string
	CumQuantity      float64
	CumQuote         float64
	ExecutedQuantity float64
	OrderId          int
	AveragePrice     float64
	OriginalQuantity float64
	Price            float64
	ReduceOnly       bool
	Side             string
	PositionSide     string
	Status           string
	StopPrice        float64
	ClosePosition    bool
	Symbol           string
	TimeInForce      string
	Type             string
	OriginalType     string
	ActivationPrice  float64
	PriceRate        float64
	UpdateTime       int64
	WorkingType      string
	PriceProtect     bool
}

type AccountEvent interface {
	eventType() string
}

type MarginPosition struct {
	Symbol            string
	PositionSide      string
	PositionAmount    float64
	MarginType        string
	IsolatedWallet    float64
	MarketPrice       float64
	UnrealizedPnL     float64
	MaintenanceMargin float64
}

type MarginCallEvent struct {
	Event              string
	EventTime          int64
	CrossWalletBalance float64
	Positions          []MarginPosition
}

func (s MarginCallEvent) eventType() string {
	return s.Event
}

type AccountUpdateBalance struct {
	Asset              string
	WalletBalance      float64
	CrossWalletBalance float64
	BalanceChange      float64
}

type AccountUpdatePosition struct {
	Symbol         string
	PositionAmount float64
	EntryPrice     float64
	Accumulated    float64
	UnrealizedPnL  float64
	MarginType     string
	IsolatedWallet float64
	PositionSide   string
}

type AccountUpdateData struct {
	UpdateType string
	Balances   []AccountUpdateBalance
	Positions  []AccountUpdatePosition
}

type AccountUpdateEvent struct {
	Event           string
	EventTime       int64
	TransactionTime int64
	UpdateData      AccountUpdateData
}

type OrderTradeData struct {
	Symbol               string
	ClientOrderId        string
	OrderSide            string
	OrderType            string
	TimeInForce          string
	Quantity             float64
	Price                float64
	AveragePrice         float64
	StopPrice            float64
	ExectutionType       string
	OrderStatus          string
	OrderId              int64
	LastFilledQuantity   float64
	AccumulatedQuantity  float64
	LastFilledPrice      float64
	CommissionAsset      string
	Commission           float64
	TradeTime            int64
	TradeId              int64
	BidsNotional         float64
	AskNotional          float64
	IsMakerSide          bool
	IsReduceOnly         bool
	StopPriceWorkingType string
	OringalOrderType     string
	PositionSide         string
	IsCloseAll           bool
	ActivationPrice      float64
	CallbackRate         float64
	RealizedProfit       float64
}

type OrderTradeUpdateEvent struct {
	Event          string
	EventTime      int64
	TransationTime int64
	OrderData      OrderTradeData
}
