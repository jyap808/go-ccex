go-ccex
==========

go-ccex is an implementation of the Ccex API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

## Import
	import "github.com/jyap808/go-ccex"
	
## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/jyap808/go-ccex"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Ccex client
	ccex := ccex.New(API_KEY, API_SECRET)

	// Get markets
	markets, err := ccex.GetMarkets()
	fmt.Println(err, markets)
}
~~~	

See ["Examples" folder for more... examples](https://github.com/jyap808/go-ccex/blob/master/examples/ccex.go)

