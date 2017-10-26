// Package Ccex is an implementation of the Biitrex API in Golang.
package ccex

import (
	"encoding/json"
	"errors"
	"strings"
)

const (
	API_BASE                   = "https://c-cex.com/t" // Ccex API endpoint
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                      // HTTP client timeout
)

// New return a instanciate ccex struct
func New(apiKey, apiSecret string) *Ccex {
	client := NewClient(apiKey, apiSecret)
	return &Ccex{client}
}

// handleErr gets JSON response from Ccex API en deal with error
func handleErr(r jsonResponse) error {
	if !r.Success {
		return errors.New(r.Message)
	}
	return nil
}

// ccex represent a ccex client
type Ccex struct {
	client *client
}

// set enable/disable http request/response dump
func (c *Ccex) SetDebug(enable bool) {
	c.client.debug = enable
}

// GetDistribution is used to get the distribution.
func (b *Ccex) GetDistribution(market string) (distribution Distribution, err error) {
	r, err := b.client.do("GET", "api_pub.html?a=GetBalanceDistribution&currencyName="+strings.ToLower(market), "", false)
	if err != nil {
		return
	}

	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}

	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &distribution)
	return

}

func (b *Ccex) GetPairs() (pairs Pairs, err error) {
	r, err := b.client.do("GET", "pairs.json", "", false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &pairs); err != nil {
		return
	}
	return
}

func (b *Ccex) GetCoinNames() (coinnames CoinNames, err error) {
	r, err := b.client.do("GET", "coinnames.json", "", false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &coinnames); err != nil {
		return
	}
	return
}
