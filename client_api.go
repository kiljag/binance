package binance

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type secType int

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned
)

type request struct {
	method     string
	endpoint   string
	query      url.Values
	recvWindow int64 // should be set in milliseconds
	secType    secType
	header     http.Header
	fullURL    string
}

func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setParams(m map[string]interface{}) *request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

func (c *Client) parseRequest(r *request) {

	fullURL := fmt.Sprintf("%s%s", c.baseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.setParam(key_RECVWINDOW, r.recvWindow)
	}
	if r.secType == secTypeSigned {
		r.setParam(key_TIMESTAMP, CurrentTimestamp())
	}

	queryString := r.query.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set(X_MBX_APIKEY, c.apiKey)
	}

	// append signature
	if r.secType == secTypeSigned {
		v := url.Values{}
		v.Set(key_SIGNATURE, computeSignature(queryString, c.secretKey))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	log.Println("full url : " + fullURL)
	r.fullURL = fullURL
	r.header = header
}

func (c *Client) callAPI(r *request) ([]byte, error) {

	if c.weightUsed > 1000 {
		log.Println("error weight limit exceeded, sleeping for 1 min, weight used : ", c.weightUsed)
		time.Sleep(1 * time.Minute)
	}

	c.parseRequest(r)
	httpClient := &http.Client{}
	req, _ := http.NewRequest(r.method, r.fullURL, nil)

	// call http api
	req.Header = r.header
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println("error in api call", err, r.fullURL)
		return []byte{}, err
	}
	defer res.Body.Close()

	// update weight used
	weightUsed := int(ParseInt(res.Header.Get("X-Mbx-Used-Weight-1m")))
	if weightUsed > 0 {
		c.weightUsed = weightUsed
	}

	// read api response
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error in reading api response : ", err)
		return []byte{}, err
	}

	// check for status code
	if res.StatusCode != 200 {
		log.Printf("error got invalid status code (%d) : %s\n", res.StatusCode, string(data))
		return []byte{}, fmt.Errorf("invalid status code(%d) : %s", res.StatusCode, string(data))
	}

	return data, nil
}
