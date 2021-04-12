爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
选出任一 x，满足 0 < x < N 且 N % x == 0 。
用 N - x 替换黑板上的数字 N 。
如果玩家无法执行这些操作，就会输掉游戏。
只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 False。假设两个玩家都以最佳状态参与游戏。
 示例 1：
输入：2
输出：true
解释：爱丽丝选择 1，鲍勃无法进行操作。
示例 2：
输入：3
输出：false
解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。
 
提示：
1 <= N <= 1000
1. Clearfication:

2. Coding:

3. 看题解：
记忆话递归，有重复子问题，因此可以使用记忆化递归
func divisorGame(n int) bool {
    if n == 1 {
        return false
    }
    
    memo := make([]int, n + 1)
    return dfs(n, memo)
}
func dfs(n int,memo[] int)bool {
    if t := memo[n]; t != 0 {
        return t == 1
    }
    if n == 1 {
        return false
    }
    
    if n == 2 {
        return true
    }
      
    canWin := -1
    
    for i := 1;i <= n / 2;i++ {
        if n % i == 0 && !dfs(n - i,memo) {
            canWin = 1
            break
        }
    }
    
    memo[n] = canWin
    return canWin == 1
}
比较巧妙的地方使用了 是否是0判断map里面这个值是否已经计算过了
动态规划：
func divisorGame(n int) bool {
    if n == 1 {
        return false
    }
    f := make([]bool,n + 1)
    f[1],f[2] = false,true
    for i := 3;i <= n;i++ {
        for j := 1;j < i;j++ {
            if i % j == 0 && !f[i - j] {
                f[i] = true
                break
            }
        }
    }
    
    return f[n]
}

4. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(n) / O(1): 记忆化数组 or 变量

5. 总结：
5.1：第一次遇到博弈问题，说实话有点懵逼
5.2: 不知道如何下手和分析
