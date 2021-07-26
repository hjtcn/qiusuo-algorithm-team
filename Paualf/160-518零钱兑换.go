给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。 
题目数据保证结果符合 32 位带符号整数。
 
示例 1：
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2：
输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3 。
示例 3：
输入：amount = 10, coins = [10] 
输出：1
 
提示：
1 <= coins.length <= 300
1 <= coins[i] <= 5000
coins 中的所有值 互不相同
0 <= amount <= 5000

1. Clearficaion:
        我理解的完全背包，每次都可以往里面放东西，如果放的下的话，没有个数的限制
        感觉用递归是可以解决的
Coding:
func change(amount int, coins []int) int {
    ret := 0
    helper(coins, amount,0,&ret)
    return ret
}
func helper(coins []int, amount int,currentAmount int,ret *int) {
    // terminator
    if currentAmount == amount {
        *ret = *ret + 1
        return
    }
    // process current logic
    for i := 0;i < len(coins);i++ {
        if coins[i] + currentAmount <= amount {
            helper(coins, amount, coins[i] + currentAmount, ret)
        }
    }
}

结果不符合预期，输出的结果为9
想了一下，没有对结果进行去重
超时了

func change(amount int, coins []int) int {
    ret := 0
    path := []int{}
    m := make(map[string]int)
    helper(coins, amount,0,&ret,m,path)
    return ret
}

func helper(coins []int, amount int,currentAmount int,ret *int,m map[string]int,path []int) {
    // terminator
    if currentAmount == amount {
        sort.Ints(path)
        key := getMapKey(path)
        if _,ok := m[key];!ok {
            m[key] = 1
            *ret = *ret + 1
        }
        
        return
    }
    // process current logic
    for i := 0;i < len(coins);i++ {
        if coins[i] + currentAmount <= amount {
            helper(coins, amount, coins[i] + currentAmount, ret,m,append(path, coins[i]))
        }
    }
}

func getMapKey(path []int) (ret string) {
    ret = ""
    for _,v := range path {
        ret = ret + strconv.Itoa(v)
    }
    return ret
}

2. 看题解：
看题解中的代码是蛮简单的，
感觉真的自己能写出来的话，里面还是有很多思考在里面的

func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            dp[i] += dp[i-coin]
        }
    }
    return dp[amount]
}

动规五部曲：
1. 确定dp数组以及下标的含义
dp[j]: 凑成总金额j的货币组合数 dp[j]

2. dp公式
dp[j]:考虑所有的coins[i] 的组合总和，就是所有的dp[j - coins[i]] 相加
所以递推公式： dp[j] += dp[j - coins[i]]

3. dp数组如何初始化
dp[0] = 1
从 dp[i] 的含义上来讲就是：凑成总金额为0的货币组合数为1

4. 确定遍历顺序
外层 for 循环遍历物品（钱币），内层for循环遍历背包（金钱总额），还是外层for遍历背包（金钱总额），内层for循环遍历物品（钱币）
这里的遍历顺序，确实有点不是很好理解
声明时候遍历的是组合数，声明时候遍历的是排列数。。。又是一个细节

5. 举例推导dp数组

3. 复杂度分析：
时间复杂度：O(amount x n ),其中 amount是总金额，n是数组coins的长度
空间复杂度：O(amount) amount 是总金额

4. 总结：
4.1： dp里面还是有蛮多细节的，要提前分析清楚，再去写代码，这个思维方式有利于我们日常工作的，像盖房子一样，只有前期把设计图做的很详细，将各种边界和可能出现的问题考虑清楚，施工的时候遇到的问题才会少一点
