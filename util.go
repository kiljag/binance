package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"time"
)

func CurrentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func computeSignature(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(message))
	if err != nil {
		log.Println("error in computing signature : ", err)
	}
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func parseFloat(str string) float64 {
	if str == "" {
		return 0
	}
	a, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println("error in parsing float : ", err)
	}
	return a
}

func ParseInt(str string) int64 {
	if str == "" {
		return 0
	}
	a, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("error in parsing integer : ", err)
	}
	return a
}
