package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}



	sort.Ints(coins)



	buffer := amount
	count := 0
	for buffer > 0 {

		if len(coins)-1 <= buffer{
			buffer -= (len(coins)-1)
			count++
			

		}else if len(coins)-1 > buffer {
			coins = coins[:(len(coins)-1)]

			count++
		}

	}



	return -1

}

func main() {
	coinChange([]int{1, 2, 5}, 11)
	coinChange([]int{2}, 3)
	coinChange([]int{1}, 0)
}
