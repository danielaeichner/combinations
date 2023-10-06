package main

import (
	"fmt"
	"slices"
)

type pair struct {
	pairAddress string
	tokens      []string
}

func reverseNum(num int) int {
	if num == 0 {
		return 1
	}
	return 0
}

func main() {

	list := []pair{
		{
			pairAddress: "0x1",
			tokens:      []string{"0x3", "0x4"},
		},
		{
			pairAddress: "0x2",

			tokens: []string{"0x3", "0x4"},
		},
	}

	for _, outer := range list {

		for j := 0; j < 2; j++ {

			for _, inner := range list {

				// do not make combination with the same pair
				if inner.pairAddress != outer.pairAddress {
					// check if the receiving token is in this pair
					if slices.Contains(inner.tokens, outer.tokens[reverseNum(j)]) {

						fmt.Println("startToken", outer.tokens[j])
						fmt.Println(outer.pairAddress, outer.tokens[0], outer.tokens[1])
						fmt.Println(inner.pairAddress, inner.tokens[0], inner.tokens[1])
						fmt.Println("---")
					}
				}

			}
		}
	}
}
