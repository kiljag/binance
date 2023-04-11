package binance

import "fmt"

type AccountService struct {
	c        *Client
	balances []CoinBalance
}

func (c *Client) NewAccountService() *AccountService {
	return &AccountService{c: c}
}

func (s *AccountService) GetBalances() []CoinBalance {
	return s.getBalances()
}

func (s *AccountService) SetLeverage(symbol string, leverage int) bool {
	return s.updateLeverage(symbol, leverage)
}

func (s *AccountService) SetMarginType(symbol string, marginType MarginType) bool {
	return s.updateMarginType(symbol, marginType)
}

func (s *AccountService) PlaceLimitOrder(info *InfoSymbol, order *LimitOrder) (*OrderResponse, error) {

	orderService := OrderService{
		c:           s.c,
		Symbol:      order.Symbol,
		Side:        order.Side,
		OrderType:   OrderTypeLimit,
		TimeInForce: order.TimeInForce,
		Quantity:    fmt.Sprintf("%.5f", order.Quantity),
		Price:       fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.Price),
	}
	return orderService.placeOrder()
}

func (s *AccountService) PlaceMarketOrder(info *InfoSymbol, order *MarketOrder) (*OrderResponse, error) {
	orderService := OrderService{
		c:         s.c,
		Symbol:    order.Symbol,
		Side:      order.Side,
		OrderType: OrderTypeMarket,
		Quantity:  fmt.Sprintf("%f", order.Quantity),
	}
	return orderService.placeOrder()
}

func (s *AccountService) PlaceStopOrder(info *InfoSymbol, order *StopOrder) (*OrderResponse, error) {
	orderService := OrderService{
		c:          s.c,
		Symbol:     order.Symbol,
		Side:       order.Side,
		OrderType:  OrderTypeStop,
		Quantity:   fmt.Sprintf("%f", order.Quantity),
		Price:      fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.Price),
		StopPrice:  fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.StopPrice),
		ReduceOnly: order.ReduceOnly,
	}
	return orderService.placeOrder()
}

func (s *AccountService) PlaceTakeProfitOrder(info *InfoSymbol, order *TakeProfitOrder) (*OrderResponse, error) {
	orderService := OrderService{
		c:          s.c,
		Symbol:     order.Symbol,
		Side:       order.Side,
		OrderType:  OrderTypeTakeProfit,
		Quantity:   fmt.Sprintf("%f", order.Quantity),
		Price:      fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.Price),
		StopPrice:  fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.StopPrice),
		ReduceOnly: order.ReduceOnly,
	}
	return orderService.placeOrder()
}

func (s *AccountService) PlaceStopMarketOrder(info *InfoSymbol, order *StopMarketOrder) (*OrderResponse, error) {
	orderService := OrderService{
		c:          s.c,
		Symbol:     order.Symbol,
		Side:       order.Side,
		OrderType:  OrderTypeStopMarket,
		Quantity:   fmt.Sprintf(fmt.Sprintf("%%.%df", info.QuantityPrecision), order.Quantity),
		StopPrice:  fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.StopPrice),
		ReduceOnly: order.ReduceOnly,
	}
	return orderService.placeOrder()
}

func (s *AccountService) PlaceTakeProfitMarketOrder(info *InfoSymbol, order *TakeProfitMarketOrder) (*OrderResponse, error) {
	orderService := OrderService{
		c:          s.c,
		Symbol:     order.Symbol,
		Side:       order.Side,
		OrderType:  OrderTypeTakeProfitMarket,
		Quantity:   fmt.Sprintf(fmt.Sprintf("%%.%df", info.QuantityPrecision), order.Quantity),
		StopPrice:  fmt.Sprintf(fmt.Sprintf("%%.%df", info.PricePrecision), order.StopPrice),
		ReduceOnly: order.ReduceOnly,
	}
	return orderService.placeOrder()
}

func (s *AccountService) CancelOrder(symbol string, orderId int64) (*OrderResponse, error) {
	orderService := OrderService{
		c: s.c,
	}
	return orderService.cancelOrder(symbol, orderId)
}

func (s *AccountService) CancelAllOpenOrders(symbol string) bool {
	orderService := OrderService{
		c: s.c,
	}
	return orderService.cancelAllOpenOrders(symbol)
}
