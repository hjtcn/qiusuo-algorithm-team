给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
示例：
输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。
 
提示：
1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100
    1. Clarification:
        两个数组中： 公共的、长度最长的子数组的长度
        1  2 3 2 1
     3  0  0 1 1 1 
     2  0  0 1 2 2
     1  1  1 1 2 3
     4  1
     7  1
    动态规划5步曲：
    脑子里面忽然想到了，不止两个数组，怎么求呢？哈哈哈，不要那么发散，先把眼前的问题给解决好

    1. dp定义
        dp[i][j]: 表示 A数组 [0-i] B数组 [0-j] 之间的最长子数组的长度

    2. dp 方程
        if A[i] == B[j]
            dp[i][j] = dp[i-1][j-1] + 1
        else 
            dp[i][j] = max(dp[i-1][j],dp[i][j-1])

    3. 初始化
        dp[0][0] = 0 : 因为没有元素的时候，没有子数组，所以长度为0
        for i := 1;i < len(A);i++ {
            if nums1[i] == nums2[0] {
                dp[0][i] = dp[0][i-1]+1
            }
        }
        for j := 1;j < len(B);j++ {
            if nums2[j] == nums1[0] {
                dp[j][0] = dp[j-1][0] + 1
            }
        }
        修改：
        run 这个case的时候一些修改
        [0,0,0,0,0]
        [0,0,0,0,0]
        for i := 1;i < len(A);i++ {
            if nums1[i] == nums2[0] {
                dp[0][i] = 1
            }else {
                dp[0][i] = dp[0][i-1]
            }
        }
        for j := 1;j < len(B);j++ {
            if nums2[j] == nums1[0] {
                dp[j][0] = 1
            }else {
                dp[j][0] = dp[j-1][0]
            }
        }
    4. 遍历顺序

    5. case by case 

func findLength(nums1 []int, nums2 []int) int {
    len1,len2 := len(nums1),len(nums2)
    if len1 == 0 || len2 == 0 {
        return 0
    }

    dp := make([][]int,len2 + 1)
    for i := 0;i < len1 + 1;i++ {
        dp[i] = make([]int, len1 + 1)
    }
    dp[0][0] = 0

   for i := 1;i < len1;i++ {
            if nums1[i] == nums2[0] {
                dp[0][i] = 1
            }else {
                dp[0][i] = dp[0][i-1]
            }
        }

       for j := 1;j < len2;j++ {
            if nums2[j] == nums1[0] {
                dp[j][0] = 1
            }else {
                dp[j][0] = dp[j-1][0]
            }
        }

     for i := 1;i <= len1;i++ {
         for j := 1;j <= len2;j++ {
             if nums1[i-1] == nums2[j-1] {
                 dp[i][j] = dp[i-1][j-1] + 1
             }else {
                 dp[i][j] = max(dp[i][j-1],dp[i-1][j])
             }
         }
     }

    return dp[len1][len2]
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

初始化和边界条件没有想清楚。。。。
脑子转不动了，晚上来了看题解

2. 看题解：

func findLength(A []int, B []int) int {
    m, n := len(A), len(B)
    res := 0
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ { 
        dp[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if A[i-1] == B[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            }
            if dp[i][j] > res {
                res = dp[i][j]
            }
        }
    }
    return res
}

上面初始化错了哈。。。。，边界条件也搞错了
看了官方题解从后遍历，也是蛮好的，但是这道题我没想到为什么会想到从后往前遍历

3. 复杂度分析：
时间复杂度：O(n x m)
空间复杂度：O(n)

4. 总结：
4.1: 初始化错了导致后面都错了哈，调试也只是乱调，所以一开始分析清楚挺重要的，但是自己有时候觉得一开始不愿意花很多时间去分析，一上来就像写代码，这样往往代码写的不咋地，所以要养成分析清楚题目的好习惯
