package binance

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type WebSocketStream struct {
	isActive   bool
	url        string
	wsConn     *websocket.Conn
	wsOpenTime int64
	timeout    int64
}

func (s *WebSocketStream) getNextMessage() ([]byte, error) {
	if !s.isActive {
		return nil, fmt.Errorf("channel is inactive")
	}

	if s.wsOpenTime == 0 || (CurrentTimestamp()-s.wsOpenTime) > s.timeout {
		if s.wsConn != nil {
			s.wsConn.Close()
		}
		log.Println("opening wstream : " + s.url)
		c, _, err := websocket.DefaultDialer.Dial(s.url, nil)
		if err != nil {
			log.Println("error in opening wstream : ", err, s.url)
			return nil, err
		}
		s.wsConn = c
		s.wsOpenTime = CurrentTimestamp()
	}

	_, msg, err := s.wsConn.ReadMessage()
	if err != nil {
		log.Println("error in reading ws message : ", err, s.url)
		return nil, err
	}
	return msg, err
}
