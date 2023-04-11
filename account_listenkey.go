package binance

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonListenKey struct {
	ListenKey string `json:"listenKey"`
}

func (s *AccountStream) getListenKey() string {
	req := request{
		method:   http.MethodPost,
		endpoint: endPointListenKey,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(&req)
	if err != nil {
		return ""
	}

	var res jsonListenKey
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println("error in parsing listen key : ", err, string(data))
		return ""
	}
	return res.ListenKey
}
