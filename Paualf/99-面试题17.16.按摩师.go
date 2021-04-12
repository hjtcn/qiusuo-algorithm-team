一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
注意：本题相对原题稍作改动
示例 1：
输入： [1,2,3,1]
输出： 4
解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
示例 2：
输入： [2,7,9,3,1]
输出： 12
解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
示例 3：
输入： [2,1,4,5,3,1,1,3]
输出： 12
解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
1. Clearfication:
每个位置都有两个状态：选择 or 不选择, 0：不选择，1:选择
总预约时间最长，限制条件是 : 不能接受相邻的预约
当前位置最大值 组成条件为 ：
选择当前值当前值 + 上一轮不能选的值：dp[i][1] + dp[i -1][0] 
不选择当前值，则上一轮是可以进行选择的: max(dp[i-1][1],dp[i-1][0])
一开始我是这样写的，好像不是很完善,运行了一下用例，有些case没有通过

2. Coding:
不完善的代码, run 起来有些case通过不了
func massage(nums []int) int {
     if len(nums) == 0 {
         return 0
     }
     n := len(nums)
     dp := make([][]int, n)
     for i := 0;i < n;i++ {
         dp[i] = make([]int, 2)
     }
     // 初始化
     dp[0][0] = 0
     dp[0][1] = 1
     for i := 1;i < n;i++ {
         dp[i][0] = dp[i - 1][1]
         dp[i][1] = dp[i - 1][0] + nums[i]
     }
   
    return max(dp[n - 1][0],dp[n - 1][1])
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

3. 看题解：
看完题解以后发现自己分析的方向是正确的，不完善的地方在于 dp[i][0] 不选择的话最大值是由前一个选择的最大值和不选择的最大值之间进行比较的，上面的代码自己初始化的地方也写错了，dp[0][1] = nums[0] 自己写成了 dp[0][1] = 1 了
func massage(nums []int) int {
     if len(nums) == 0 {
         return 0
     }
     n := len(nums)
     dp := make([][]int, n)
     for i := 0;i < n;i++ {
         dp[i] = make([]int, 2)
     }
     // 初始化
     dp[0][0] = 0
     dp[0][1] = nums[0]
     for i := 1;i < n;i++ {
         dp[i][0] = max(dp[i - 1][1],dp[i - 1][0])
         dp[i][1] = dp[i - 1][0] + nums[i]
     }
   
    return max(dp[n - 1][0],dp[n - 1][1])
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}
只用变量进行存储计算
func massage(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    dp0,dp1 := 0,nums[0]
    for i := 1;i < n;i++ {
        tdp0 := max(dp0, dp1)
        tdp1 := dp0 + nums[i]
        dp0 = tdp0
        dp1 = tdp1
    }
    return max(dp0, dp1)
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 或者 O(1)

5. 总结：
5.1: 动态规划的题目感觉对分析能力要求比较高
5.2: 分析清楚问题定义，将问题分析和理解清楚，初始化后推导动态转移方程
