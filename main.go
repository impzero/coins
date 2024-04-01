package main

import (
	"fmt"
)

func main() {

	oneCoin := NewCoin(1)
	threeCoin := NewCoin(3)
	fourCoin := NewCoin(4)
	sixCoin := NewCoin(6)

	allCoins := []CoinChain{oneCoin, threeCoin, fourCoin, sixCoin}
	cs := sixCoin.GenChain(allCoins)
	nums := []Coin{}
	for _, c := range cs {
		nums = append(nums, c.Value)
	}

	fmt.Printf("You can use %v to form a %d", nums, sixCoin.Value)
}

type Coin int

// coins in coinsconstructor shouldnt be bigger then Value because CoinChain is a node in a tree, where the root is
// the largest value coin
// should be refactored
// right now we don't have the knowledge of what coins make up other coins
// 3 could be made out of 1 and 3 but we have to always generate it which is unefficient
type CoinChain struct {
	Value           Coin
	CoinsContructor func(coins []CoinChain) []CoinChain
}

// create a tree ctor
// root (val) and coins (nodes)

// 6 - 1, 3, 4
// 6-1 > 0 yes -> 1; 5-3 yes 2 - 4 no
// 6-3 yes 3 - 4 no
// 6 - 4 yes

func (c CoinChain) GenChain(allCoins []CoinChain) []CoinChain {
	return c.CoinsContructor(allCoins)
}

func (c CoinChain) NextCoin(allCoins []CoinChain) CoinChain {
	coins := c.GenChain(allCoins)
	return coins[len(coins)-1]
}

func CoinsConstructor(val Coin) func(coins []CoinChain) []CoinChain {
	return func(coins []CoinChain) []CoinChain {
		stopIndex := 1
		for i, c := range coins {
			if c.Value < val {
				stopIndex = i + 1
			}
		}

		constructs := []CoinChain{}
		for _, c := range coins[:stopIndex] {
			if val%c.Value == 0 {
				constructs = append(constructs, c)
			} else {
				for _, construct := range constructs {
					if c.Value%construct.Value == 0 {
						constructs = append(constructs, c)
					}
					CoinFormations
				}
			}
		}
		return constructs
	}
}

func NewCoin(val Coin) CoinChain {
	return CoinChain{
		Value:           val,
		CoinsContructor: CoinsConstructor(val),
	}
}
