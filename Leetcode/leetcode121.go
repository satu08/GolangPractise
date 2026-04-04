package main

//	func main() {
//		fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
//	}
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
