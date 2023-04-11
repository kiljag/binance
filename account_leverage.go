package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonLeverage struct {
	Leverage       int    `json:"leverage"`
	MaxNotionValue string `json:"maxNotionValue"`
	Symbol         string `json:"symbol"`
}

func (s *AccountService) updateLeverage(symbol string, leverage int) bool {
	req := request{
		method:   http.MethodPost,
		endpoint: endPointLeverage,
		secType:  secTypeSigned,
	}
	req.setParam(key_SYMBOL, symbol)
	req.setParam(key_LEVERAGE, leverage)

	data, err := s.c.callAPI(&req)
	if err != nil {
		return false
	}

	var res jsonLeverage
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println("error in parsing leverage json : ", err, string(data))
		return false
	}
	return res.Leverage == leverage
}
