package main

import "fmt"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	return -1

}

func main() {
	fmt.Println("test")

	var amount int = 1

	coinChange([]int{1, 2, 5}, 11)
}
