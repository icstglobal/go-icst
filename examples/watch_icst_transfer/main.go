package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"

	"github.com/icstglobal/go-icst/chain"

	"github.com/icstglobal/go-icst/chain/eth"
)

var start = flag.Int64("start", -1, "block number to start from")
var url = flag.String("url", "", "url of the ethereum node")

func main() {
	flag.Parse()

	if len(*url) == 0 {
		flag.Usage()
		return
	}
	ether, err := eth.DialEthereum(*url)
	if err != nil {
		fmt.Print(err)
		return
	}
	chain.Set(chain.Eth, ether)

	blc, err := chain.Get(chain.Eth)
	if err != nil {
		fmt.Print(err)
	}

	var blockStart *big.Int
	if *start != -1 {
		blockStart = big.NewInt(*start)
	}
	blocks, errors := blc.WatchICSTTransfer(context.Background(), blockStart)
	for {
		select {
		case block := <-blocks:
			fmt.Printf("block:%+v\n", block)
			for _, trans := range block.Trans {
				fmt.Printf("trans: %v\n", trans)
			}
		case err := <-errors:
			fmt.Print(err)
		}
	}
}
