func maxProfit(prices []int) int {
    maxProfit := 0
    buy := prices[0]

    for i := 1 ; i <= len(prices) - 1 ; i++ {
        if prices[i] < buy {
            buy = prices[i]
        }
        if prices[i] - buy > maxProfit {
            maxProfit = prices[i] - buy
        }
    }
    return maxProfit
}
