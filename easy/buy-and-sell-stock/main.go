package main

import "fmt"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
//	largest = max(largest, prices[i]+prices[i-1]-prices[i-1])
func maxProfit(prices []int) int {
	var profit int
	smallest := prices[0]
	for i := 1; i < len(prices); i++ {
		smallest = min(smallest, prices[i])
		profit = max(profit, prices[i]-smallest)
	}
	return profit
}

func main() {
	//	prices := []int{7, 1, 5, 3, 6, 4}
	//	prices := []int{1, 2}
	prices := []int{7, 6, 5, 4, 3, 2}
	fmt.Println(maxProfit(prices))
}
