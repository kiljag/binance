
# Binance Futures API (Golang)


### Initialization

```golang

import (
    "os"
    "github.com/kiljag/binance"
)

// initialize binance client
apiKey := os.Getenv("BINANCE_API_KEY")
secretKey := os.Getenv("BINANCE_SECRET_KEY")
client := binance.NewClient(apiKey, secretKey)

```

### Exchange Info
To get all active symbols and other symbol related details
``` golang

client.GetExchangeInfo() // returns *ExchangeInfo

```

### Klines Stream
To stream real time market klines for a symbol

```golang

symbol := "BTCUSDT"
interval := "15m"
dropProb:= 0.15 // to get klines less frequently
klineStream := client.NewKlineStream(symbol, interval, dropProb)
klinesChannel := klinesStream.Start()

for k, ok := <- klinesChannel {
    if !ok {
        break
    }
    // process k
}

```


### Account Service
1. Placing a trade (LIMIT, MARKET, etc..)


```golang

accountService := client.NewAccountService()

// 1. To get available balances of each coin
func (s *AccountService) GetBalances() []CoinBalance

// 2. To set leverage for a particular symbol
func (s *AccountService) SetLeverage(symbol string, leverage int) bool

// 3. to set margin type
func (s *AccountService) SetMarginType(symbol string, marginType MarginType) bool

// 4. to place different types of orders
func (s *AccountService) PlaceLimitOrder(info *InfoSymbol, order *LimitOrder) (..)
func (s *AccountService) PlaceMarketOrder(info *InfoSymbol, order *MarketOrder) (..)
func (s *AccountService) PlaceStopOrder(info *InfoSymbol, order *StopOrder) (..)
func (s *AccountService) PlaceTakeProfitOrder(info *InfoSymbol, order *TakeProfitOrder) (..)
func (s *AccountService) PlaceStopMarketOrder(info *InfoSymbol, order *StopMarketOrder) (..)
func (s *AccountService) PlaceTakeProfitMarketOrder(info *InfoSymbol, order *TakeProfitMarketOrder) (..)

// 5. to cancel orders
func (s *AccountService) CancelOrder(symbol string, orderId int64) (..)
func (s *AccountService) CancelAllOpenOrders(symbol string) bool

```