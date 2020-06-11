/**
 *	题目：买卖股票的最佳时机
 */
package go_code

import(
	"testing"
)

func maxProfit(prices []int) int {
	Max := 0
	for i := 0; i < len(prices) - 1; i++{
		for j := i + 1; j < len(prices); j++ {
			delta := prices[j] - prices[i]
			if delta > Max {
				Max = delta
			}
		}
	} 
	return Max
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	Min := prices[0]
	Max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] - Min > Max {
			Max = prices[i] - Min
		}
		if prices[i] < Min {
			Min = prices[i]
		}
	}
	return Max
}

func TestBestTime(t *testing.T) {
	prices := []int{7,6,4,3,1}
	t.Log(maxProfit2(prices))
}