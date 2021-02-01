/*
泰波那契序列 Tn 定义如下： 
T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
示例 1：
输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
示例 2：
输入：n = 25
输出：138953
 
提示：
0 <= n <= 37
答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
*/
Tn+3 = T(n) + T(n + 1) + T(n + 2)

1. Clearfication:
n == 0 => 0
n == 1 => 1
n == 2 => 1
n = 3  => T(n-1) + T(n - 2) + T(n-3)
    => T(2) + T(1) + T(0)
递归解决 =》 超时
为了避免超时 =》加一个记忆数组进行处理
最后是动态规划
分解子问题：
伪代码：
if n == 0 
    return 0
if n == 1
    return 1
if n == 2 
    return 1
return f(n - 1) + f (n - 2) + f (n - 3)

Coding:
func tribonacci(n int) int {
     if n == 0 {
         return 0
     }
     if n == 1 {
         return 1
     }
     if n == 2 {
         return 1
     }
     return tribonacci(n - 1) + tribonacci(n - 2) + tribonacci(n - 3)
}

超时：
使用 记忆化数组加速访问：
1. 数组加到哪里？
2. 怎么返回结果

伪代码：
memo := make([]int, n)
// 使用记忆化数组加速内容访问
int helper(n int,memo *int) {
    
    if memo[n] > 0 {
        return memo[n]
    }
    
    memo[n] = helper(n-1,memo) + helper(n-2,memo) + helper(n-3,memo)
    
    return memo[n]
} 

想了一下下面这种边界条件判断，是可以放在调用函数层或者是被调用的函数里面的
可以类比到工作中我们使用到的Api 和 Service的调用关系，我觉得两者都需要加上条件的判断，这样方便程序更好的处理

if n == 0 {
    return 0
}
if n == 1 {
    return 1
}
if n == 2 {
  return 1
}    

Coding:
func tribonacci(n int) int {
     if n == 0 {
         return 0
     }
     if n == 1 {
         return 1
     }
     if n == 2 {
         return 1
     }
    memo := make([]int, n + 1)
    
    return helper(n, &memo)
}

func helper(n int,memo *int) int {
    
    if memo[n] > 0 {
        return memo[n]
    }
    
    memo[n] = helper(n-1,memo) + helper(n-2,memo) + helper(n-3,memo)
    
    return memo[n]
} 

修改后的版本：上面代码中的 memo *[]int ，指向切片的指针也错了，下次要再注意下
func tribonacci(n int) int {
     if n == 0 {
         return 0
     }
     if n == 1 {
         return 1
     }
     if n == 2 {
         return 1
     }
    memo := make([]int, n + 1)
    
    return helper(n, &memo)
}
func helper(n int,memo *[]int) int {
    if n == 0 {
         return 0
     }
     if n == 1 {
         return 1
     }
     if n == 2 {
         return 1
     }
    if (*memo)[n] > 0 {
        return (*memo)[n]
    }
    
    (*memo)[n] = helper(n-1,memo) + helper(n-2,memo) + helper(n-3,memo)
    
    return (*memo)[n]
} 

看到里面有两段重复的判断，我是这样想的，分离关注点，Api的判断和Service 的判断是都要加的，所以就有了重复的判断逻辑
动态规划怎么写呢？

1. 开辟一个数组用于存储每个结果
2. 动态递推
伪代码：
dp := make([]int, n + 1)
dp[0] = 0
dp[1]  = 1
dp[2] = 1
dp[n] = dp[n-1] + dp[n-2] + dp[n-3]
return dp[n]

Coding:
func tribonacci(n int) int {
     if n == 0 {
         return 0
     }
     if n == 1 {
         return 1
     }
     if n == 2 {
         return 1
     }
    
     dp := make([]int, n + 1)
     dp[0] = 0
     dp[1] = 1
     dp[2] = 1
    
    for i := 3;i <= n;i++ {
        dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3]
    }
    
    return dp[n]
}

3. 看题解：
使用map定义和记录还是蛮好玩的

func tribonacci(n int) int {
    cache := map[int]int{
        0 : 0, 1 : 1, 2 : 1,
    }
    return recursive(cache, n)
}
func recursive(cache map[int]int, n int) int {
    if _, ok := cache[n]; !ok {
        cache[n] = recursive(cache, n-3) + recursive(cache, n-2) + recursive(cache, n-1)
    }
    return cache[n]
}

Good Code:
func tribonacci(n int) int {
    seq := []int{0,1,1}
    if n <= 2 {
        return seq[n]
    }
    
    for i := 2; i < n; i++ {
        seq[0], seq[1], seq[2] = seq[1], seq[2], seq[0]+seq[1]+seq[2]
    }
    return seq[2]
}

4. 复杂度分析

时间复杂度：
递归：O(3^n) 阶乘的时间复杂度
记忆话数组：O(n)
动态规划：O(n)

空间复杂度：
    递归：开辟栈空间 O(3^n)
    记忆话数组：O(n) 开辟数组或者map存储节点关系
    动态规划：O(n) 开辟数组存储节点

5. 总结:
5.1. 设计代码组织结构的时候可以考虑分离关注点，将逻辑拆分出去

5.2. 递归的时间复杂度要好好看下

5.3. 一定要多看题解哈，如果不看题解的话，这道题做不做没有区别的哈
