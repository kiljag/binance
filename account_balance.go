package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonCoinBalance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"`
	UpdateTime         int64  `json:"updateTime"`
}

func (s *AccountService) getBalances() []CoinBalance {

	req := request{
		method:     http.MethodGet,
		endpoint:   endPointBalance,
		recvWindow: 5000,
		query:      nil,
		secType:    secTypeSigned,
	}

	body, err := s.c.callAPI(&req)
	if err != nil {
		return []CoinBalance{}
	}

	var balanceObjList []jsonCoinBalance
	err = json.Unmarshal(body, &balanceObjList)
	if err != nil {
		log.Println("error in parsing balance json : ", err, string(body))
		return []CoinBalance{}
	}

	s.balances = make([]CoinBalance, 0)
	for _, obj := range balanceObjList {
		balance := CoinBalance{
			AccountAlias:       obj.AccountAlias,
			Asset:              obj.Asset,
			Balance:            parseFloat(obj.Balance),
			CrossWalletBalance: parseFloat(obj.CrossWalletBalance),
			CrossUnPnl:         parseFloat(obj.CrossUnPnl),
			AvailableBalance:   parseFloat(obj.AvailableBalance),
			MaxWithdrawAmount:  parseFloat(obj.MaxWithdrawAmount),
			MarginAvailable:    obj.MarginAvailable,
			UpdateTime:         obj.UpdateTime,
		}
		s.balances = append(s.balances, balance)
	}
	return s.balances
}
