package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	buffer := amount

	if buffer == 0 {
		return 0
	}
	if len(coins) == 1 {
		buffer = coins[0] - buffer
		return buffer
	}

	sort.Ints(coins)

	count := 0
	for buffer >= 0 && len(coins) > 0 {

		if coins[(len(coins)-1)] <= buffer {
			buffer -= coins[(len(coins) - 1)]
			count++

		} else if coins[(len(coins)-1)] > buffer {
			coins = coins[:(len(coins) - 1)]

		}

	}

	return count

}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}
