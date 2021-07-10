给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
示例 1：
输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
示例 2：
输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
 
提示：
1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100

1. Clearfication:
    01背包挺像的，还是不是很会
 如果是我做的话，会使用递归
coding:

func findMaxForm(strs []string, m int, n int) int {
    ret := 0
    helper(strs,&ret,m,n,0,0)
    return ret
}

func helper(strs []string,ret *int,m,n int,current int,index int) {
    // terminator
    if m < 0 || n < 0 {
        return
    }
    if index >= len(strs) {
         if m >= 0 && n >= 0 {
             if current > *ret {
                 *ret = current
             }
         }
         return
    }    
    // process current logic
    // 不放strs[index]
    helper(strs,ret,m,n,current,index+1)
    // 放入strs[index]
    for _,v := range strs[index]{
        if v == '0' {
            m--
        }
        if v == '1' {
            n--
        }
    }
    if m >= 0 && n >=0 {
         current = current + 1
    }
   
    helper(strs,ret,m,n,current,index + 1)
    // dirll down
}

超时了，容易错的细节地方是 需要判断  if m >= 0 && n >= 0 { 而不是 m == 0 && n ==0 ，一开始的时候我是这样写的，后来调试的时候发现不太对
经典的背包问题有两个维度分别是物品和容量，有两种容量的情况下，需要使用三维动态规划求解，三个维度分别是字符串，0的容量，1的容量
定义三维数组dp，其中dp[i][j][k] 表示在前i个字符串中，使用j个0和k个1的情况下最多可以得到的字符串数量。假设数组str的长度为l,则最终答案为dp[l][m][n]

func findMaxForm(strs []string, m, n int) int {
    length := len(strs)
    dp := make([][][]int, length+1)
    for i := range dp {
        dp[i] = make([][]int, m+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, n+1)
        }
    }
    for i, s := range strs {
        zeros := strings.Count(s, "0")
        ones := len(s) - zeros
        for j := 0; j <= m; j++ {
            for k := 0; k <= n; k++ {
                dp[i+1][j][k] = dp[i][j][k]
                if j >= zeros && k >= ones {
                    dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
                }
            }
        }
    }
    return dp[length][m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

看题解分析的还是蛮详细的，现在的我知道这样想了，但是也不一定写的这么详细

3. 复杂度分析：
时间复杂度：O( l * m  * n)
空间复杂度：O(l * m *n)

4. 总结：
4.1 三维的动态规划还是蛮经典的，细节还是蛮多的

4.2 学习像题解中那样将问题描述的清晰和准确
