package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonMarginTypeResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (s *AccountService) updateMarginType(symbol string, marginType MarginType) bool {
	req := request{
		method:   http.MethodPost,
		endpoint: endPointMarginType,
		secType:  secTypeSigned,
	}
	req.setParam(key_SYMBOL, symbol)
	req.setParam(key_MARGIN_TYPE, marginType)

	data, err := s.c.callAPI(&req)
	if err != nil {
		return false
	}

	var res jsonMarginTypeResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println("error in parsing margin type response json : ", err, string(data))
		return false
	}
	return res.Code == 200
}
