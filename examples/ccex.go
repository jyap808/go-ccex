package main

import (
	"fmt"
	"github.com/jyap808/go-ccex"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Ccex client
	ccex := ccex.New(API_KEY, API_SECRET)

	// Get Candle ( OHLCV )
	/*
		markets, err := ccex.GetHisCandles("BTC-LTC", "hour")
		fmt.Println(markets, err)
	*/

	// Get markets
	/*
		markets, err := ccex.GetMarkets()
		fmt.Println(err, markets)
	*/

	// Get Ticker (BTC-VTC)
	/*
		ticker, err := ccex.GetTicker("BTC-DRK")
		fmt.Println(err, ticker)
	*/

    // Get Distribution (JBS)
	distribution, err := ccex.GetDistribution("SPHR")
    fmt.Println(err)
	for _, balance := range distribution.Distribution {
		fmt.Println(balance.BalanceD)
	}

	// Get market summaries
	/*
		marketSummaries, err := ccex.GetMarketSummaries()
		fmt.Println(err, marketSummaries)
	*/

	// Get orders book
	/*
		orderBook, err := ccex.GetOrderBook("BTC-DRK", "both", 100)
		fmt.Println(err, orderBook)
	*/

	// Market history
	/*
		marketHistory, err := ccex.GetMarketHistory("BTC-DRK", 100)
		for _, trade := range marketHistory {
			fmt.Println(err, trade.Timestamp.String(), trade.Quantity, trade.Price)
		}
	*/

	// Market

	// BuyLimit
	/*
		uuid, err := ccex.BuyLimit("BTC-DOGE", 1000, 0.00000102)
		fmt.Println(err, uuid)
	*/

	// BuyMarket
	/*
		uuid, err := ccex.BuyLimit("BTC-DOGE", 1000)
		fmt.Println(err, uuid)
	*/

	// Sell limit
	/*
		uuid, err := ccex.SellLimit("BTC-DOGE", 1000, 0.00000115)
		fmt.Println(err, uuid)
	*/

	// Cancel Order
	/*
		err := ccex.CancelOrder("e3b4b704-2aca-4b8c-8272-50fada7de474")
		fmt.Println(err)
	*/

	// Get open orders
	/*
		orders, err := ccex.GetOpenOrders("BTC-DOGE")
		fmt.Println(err, orders)
	*/

	// Account
	// Get balances
	/*
		balances, err := ccex.GetBalances()
		fmt.Println(err, balances)
	*/

	// Get balance
	/*
		balance, err := ccex.GetBalance("DOGE")
		fmt.Println(err, balance)
	*/

	// Get address
	/*
		address, err := ccex.GetDepositAddress("QBC")
		fmt.Println(err, address)
	*/

	// WithDraw
	/*
		whitdrawUuid, err := ccex.Withdraw("QYQeWgSnxwtTuW744z7Bs1xsgszWaFueQc", "QBC", 1.1)
		fmt.Println(err, whitdrawUuid)
	*/

	// Get order history
	/*
		orderHistory, err := ccex.GetOrderHistory("BTC-DOGE", 10)
		fmt.Println(err, orderHistory)
	*/

	// Get getwithdrawal history
	/*
		withdrawalHistory, err := ccex.GetWithdrawalHistory("all", 0)
		fmt.Println(err, withdrawalHistory)
	*/

	// Get deposit history
	/*
		deposits, err := ccex.GetDepositHistory("all", 0)
		fmt.Println(err, deposits)
	*/

}
