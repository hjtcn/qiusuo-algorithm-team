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
    1. 最大 子集
    2. m 个 0 和 n 个 1
    关键字：子集，m个0，n个1
    构造出所有子集
    然后判断是否满足 m 个 0 n个1 的判断条件，如果满足，计算该子集的大小，返回最大的值
    怎么构造出所有子集呢？

2. Coding:
/*
    1. 最大 子集
    2. m 个 0 和 n 个 1
    关键字：子集，m个0，n个1
    构造出所有子集
    然后判断是否满足 m 个 0 n个1 的判断条件，如果满足，计算该子集的大小，返回最大的值
    怎么构造出所有子集呢？
*/
func findMaxForm(strs []string, m int, n int) int {
    ret := 0
    paths := [][]string{}
    path := []string{}
    num := len(strs)
    generate(0,num,strs,path,&paths)
    for _,v := range paths {
        // 判断 v 里面 0的个数和1的个数
        zeroNum := 0
        oneNum := 0
        for _,v1 := range v {
            for s := range v1 {
                if s == '0' {
                    zeroNum++
                }
                if s == '1' {
                    oneNum++
                }
            }
        }
        if m == zeroNum && n == oneNum {
            size := len(v)
            if size > ret {
                ret = size
            }
        }
    }
    return ret
}

func generate(level,max int,strs []string,path []string,paths *[][]string) {
    // terminator
    if level >= max {
        tmp := make([]string, len(path))
        copy(tmp, path)
        *paths = append(*paths, tmp)
        return
    }
    // process current logic
    generate(level+1,max,strs,path,paths)
    path = append(path,strs[level])
    generate(level + 1,max,strs,path,paths)
    // drill down
    // reverse current status
    path = path[:len(path) - 1]
}

有case 没有跑过去
没有跑过去的case是
["10","0001","111001","1","0"] 3 4
调试了一下，发现是题没有看清楚，题目中给的是该子集中最多有m个0和n个1，自己的代码写成了 
m == zeroNum && n == oneNum
然后改成了 
zeroNum <=m && oneNum <= n 
超时了。。。

3. 看题解：
func findMaxForm(strs []string,m int,n int) int {
    statistic := func(str string) (int,int) {
        var zero,one = 0,0
        
        for _,char := range str {
            if char == '0' {
                zero++
            }else {
                one++
            }
        }
        
        return zero,one
    }
    
    length := len(strs)
    dp := make([][][]int, length + 1)
    
    for i := 0;i <= length;i++ {
        dp[i] = make([][]int, m + 1)
        for j := 0;j <= m;j++ {
            dp[i][j] = make([]int, n + 1)
        }
    }
    
    max := func(a,b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    for i := 1;i <= length;i++ {
        zero,one := statistic(strs[i-1])
        for j := 0;j <= m;j++ {
            for k := 0;k <= n;k++ {
                dp[i][j][k] = dp[i-1][j][k]
                
                if j >= zero && k >= one {
                    dp[i][j][k] = max(dp[i - 1][j][k], dp[i - 1][j-zero][k-one] + 1)
                }
            }
        }
    }
    
    return dp[length][m][n]
}

4. 复杂度分析：
时间复杂度：O(mnl)
空间复杂度：O(mn)

5. 总结：
5.1: 状态多的话就比较复杂，不好理解，说实话这道题还是不是很理解。。。
