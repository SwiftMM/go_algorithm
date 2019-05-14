package max_profit

import "fmt"

func MaxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	maxProfit := 0

	for _, item := range prices {
		fmt.Println("max profit", item)
		if item < min {
			min = item
		} else if item-min > maxProfit {
			maxProfit = item - min
		}
	}
	return maxProfit
}
