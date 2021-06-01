给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 
示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 
提示：
1 <= prices.length <= 105
0 <= prices[i] <= 104

1. Clearfication:
        买一次，卖一次，算差值，i:0->len(nums),j:i+1->len(nums),prices[j] - prices[i] 最大

2. Coding:
超时了
func maxProfit(prices []int) int {
    n := len(prices)
    
    if n <= 0 {
        return 0
    }
    ret := 0
    for i := 0;i < n;i++ {
        for j := i + 1;j < n;j++ {
            if prices[j] - prices[i] > ret {
                ret = prices[j] - prices[i]
            }
        }
    }
    return ret
}

3. 看题解：
标记历史最低点
func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxprofit := 0
    
    for i := 0; i < len(prices);i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        }else if prices[i] - minPrice > maxprofit {
            maxprofit = prices[i] - minPrice
        }
    }
    
    return maxprofit
}

感觉又体现了数学的思想：
Key observation: prices[2] - prices[0] = prices[2] - prices[1] + prices[1] - prices[0]
func maxProfit(prices []int) int {
    tmp := 0
    max := 0
    for i := 1; i < len(prices); i++ {
        tmp += prices[i] - prices[i-1]
        if tmp < 0 {
            tmp = 0
        }
        if tmp > max {
            max = tmp
        }
    }
    return max
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

5. 总结：
5.1: 标记最低位置
5.2: 思维不够灵活
