package binance

/* place and cancel orders */
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type OrderService struct {
	c *Client
	// order parameters
	Symbol           string
	Side             SideType
	PositionSide     PositionSideType
	OrderType        OrderType
	TimeInForce      TimeInForceType
	Quantity         string
	ReduceOnly       bool
	Price            string
	NewClientOrderId string
	StopPrice        string
	ClosePosition    bool
	ActivationPrice  string
	CallbackRate     string
	WorkingType      WorkingType
	PriceProtect     bool
	NewOrderRespType NewOrderRespType
	RecvWindow       string
}

/* places an order and returns response with order it
 * Mandatory fields for each order type
 *
 *    Type								Mandatory Parameters
 * *  LIMIT								timeInForce, quantity, price
 * *  MARKET							quantity
 * *  STOP/TAKE_PROFIT					quantity, price, stopPrice
 * *  STOP_MARKET/TAKE_PROFIT_MARKET	stopPrice
 * * TRAILING_STOP_MARKET				callbackRate
 **/
func (order *OrderService) placeOrder() (*OrderResponse, error) {
	req := request{
		method:   http.MethodPost,
		endpoint: endPointOrder,
		secType:  secTypeSigned,
	}

	params := make(map[string]interface{})
	params["symbol"] = order.Symbol
	params["side"] = order.Side
	params["type"] = order.OrderType

	if order.Quantity != "" {
		params["quantity"] = order.Quantity
	}
	if order.ReduceOnly {
		params["reduceOnly"] = order.ReduceOnly
	}
	if order.NewOrderRespType != "" {
		params["newOrderRespType"] = order.NewOrderRespType
	}
	if order.PositionSide != "" {
		params["positionSide"] = order.PositionSide
	}
	if order.TimeInForce != "" {
		params["timeInForce"] = order.TimeInForce
	}
	if order.ReduceOnly {
		params["reduceOnly"] = order.ReduceOnly
	}
	if order.Price != "" {
		params["price"] = order.Price
	}
	if order.NewClientOrderId != "" {
		params["newClientOrderId"] = order.NewClientOrderId
	}
	if order.StopPrice != "" {
		params["stopPrice"] = order.StopPrice
	}
	if order.WorkingType != "" {
		params["workingType"] = order.WorkingType
	}
	if order.PriceProtect {
		params["priceProtect"] = order.PriceProtect
	}
	if order.ActivationPrice != "" {
		params["activationPrice"] = order.ActivationPrice
	}
	if order.CallbackRate != "" {
		params["callbackRate"] = order.CallbackRate
	}
	if order.ClosePosition {
		params["closePosition"] = order.ClosePosition
	}
	if order.RecvWindow == "" {
		order.RecvWindow = "2000" // 2 seconds by default
	}

	req.setParams(params)
	recvWindow, _ := strconv.ParseInt(order.RecvWindow, 10, 64)
	if recvWindow == 0 {
		req.recvWindow = 2000
	}
	data, err := order.c.callAPI(&req)
	if err != nil {
		log.Println("error in placing order : ", err, order, string(data))
		return nil, err
	}
	return order.parseOrderResponse(data), nil
}

func (order *OrderService) cancelOrder(symbol string, orderId int64) (*OrderResponse, error) {
	req := request{
		method:   http.MethodDelete,
		endpoint: endPointOrder,
		secType:  secTypeSigned,
	}

	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["orderId"] = orderId
	req.setParams(params)
	req.recvWindow = 2000
	data, err := order.c.callAPI(&req)
	if err != nil {
		log.Println("error in cancelling order : ", err, order, string(data))
		return nil, err
	}

	return order.parseOrderResponse(data), nil
}

type jsonOrderResponse struct {
	ClientOrderId    string `json:"clientOrderId"`
	CumQuantity      string `json:"cumQty"`
	CumQuote         string `json:"cumQuote"`
	ExecutedQuantity string `json:"executedQty"`
	OrderId          int    `json:"orderId"`
	AveragePrice     string `json:"avgPrice"`
	OriginalQuantity string `json:"origQty"`
	Price            string `json:"price"`
	ReduceOnly       bool   `json:"reduceOnly"`
	Side             string `json:"side"`
	PositionSide     string `json:"positionSide"`
	Status           string `json:"status"`
	StopPrice        string `json:"stopPrice"`
	ClosePosition    bool   `json:"closePosition"`
	Symbol           string `json:"symbol"`
	TimeInForce      string `json:"timeInForce"`
	Type             string `json:"type"`
	OriginalType     string `json:"origType"`
	ActivationPrice  string `json:"activatePrice"`
	PriceRate        string `json:"priceRate"`
	UpdateTime       int64  `json:"updateTime"`
	WorkingType      string `json:"workingType"`
	PriceProtect     bool   `json:"priceProtect"`
}

func (order *OrderService) parseOrderResponse(data []byte) *OrderResponse {
	var res jsonOrderResponse
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Println("error in parsing order response : ", err, string(data))
		return nil
	}

	return &OrderResponse{
		ClientOrderId:    res.ClientOrderId,
		CumQuantity:      parseFloat(res.CumQuantity),
		CumQuote:         parseFloat(res.CumQuote),
		ExecutedQuantity: parseFloat(res.ExecutedQuantity),
		OrderId:          res.OrderId,
		AveragePrice:     parseFloat(res.AveragePrice),
		OriginalQuantity: parseFloat(res.OriginalQuantity),
		Price:            parseFloat(res.Price),
		ReduceOnly:       res.ReduceOnly,
		Side:             res.Side,
		PositionSide:     res.PositionSide,
		Status:           res.Status,
		StopPrice:        parseFloat(res.StopPrice),
		ClosePosition:    res.ClosePosition,
		Symbol:           res.Symbol,
		TimeInForce:      res.TimeInForce,
		Type:             res.Type,
		OriginalType:     res.OriginalType,
		ActivationPrice:  parseFloat(res.ActivationPrice),
		PriceRate:        parseFloat(res.PriceRate),
		UpdateTime:       res.UpdateTime,
		WorkingType:      res.WorkingType,
		PriceProtect:     res.PriceProtect,
	}
}

type jsonCancelAllOrders struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (order *OrderService) cancelAllOpenOrders(symbol string) bool {
	req := request{
		method:   http.MethodDelete,
		endpoint: endPointAllOpenOrders,
		secType:  secTypeSigned,
	}
	req.setParam("symbol", symbol)
	req.recvWindow = 2000
	data, err := order.c.callAPI(&req)
	if err != nil {
		return false
	}

	var res jsonCancelAllOrders
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println("error in parsing cancelAll response : ", err, string(data))
		return false
	}
	return res.Code == 200
}
