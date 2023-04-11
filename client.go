package binance

import (
	"log"
	"net/http"
)

const (
	// rest
	baseApiMainURL    = "https://fapi.binance.com"
	baseApiTestnetURL = "https://testnet.binancefuture.com"

	// ws
	baseWsMainURL          = "wss://fstream.binance.com/ws"
	baseWsTestnetURL       = "wss://stream.binancefuture.com/ws"
	baseCombinedMainURL    = "wss://fstream.binance.com/stream?streams="
	baseCombinedTestnetURL = "wss://stream.binancefuture.com/stream?streams="

	// endpoints
	endPointServerTime       = "/fapi/v1/time"
	endPointExchangeInfo     = "/fapi/v1/exchangeInfo"
	endPoint24hrTicker       = "/fapi/v1/ticker/24hr"
	endPointKlines           = "/fapi/v1/klines"
	endPointContinuousKlines = "/fapi/v1/continuousKlines"

	// userdata
	endPointBalance       = "/fapi/v2/balance"
	endPointAccount       = "/fapi/v2/account"
	endPointOrder         = "/fapi/v1/order"
	endPointAllOpenOrders = "/fapi/v1/allOpenOrders"
	endPointLeverage      = "/fapi/v1/leverage"
	endPointMarginType    = "/fapi/v1/marginType"
	endPointListenKey     = "/fapi/v1/listenKey"
)

type Client struct {
	apiKey     string
	secretKey  string
	baseURL    string
	debug      bool
	weightUsed int
}

func NewClient(apiKey, secretKey string) *Client {
	c := &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
		baseURL:   baseApiMainURL,
	}
	return c
}

func (c *Client) UseTestNet() {
	c.baseURL = baseApiTestnetURL
}

func (c *Client) DebugMode() {
	c.debug = true
}

func (c *Client) PrintServerTime() {
	req := request{
		method:   http.MethodGet,
		endpoint: endPointServerTime,
	}

	data, err := c.callAPI(&req)
	if err != nil {
		log.Println("error in reading server time : ", err)
	}
	log.Println("server time :", string(data))
}
