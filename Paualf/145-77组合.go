给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
示例:
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

1. Clearfication:
分解子问题：
        n个盒子，每个盒子都有自己的编号 1 - n 你可以选或者不选，k = 2
        terminator
            path = 2
        process current logic
        drill down
    迭代的思路使用for循环去进行遍历求解，迭代的不是很熟悉
    for i := index;i < len(num);i++ {
    }

2. Coding:
func combine(n int, k int) [][]int {
    paths := [][]int{}
    path := []int{}
    if n <= 0 || k <= 0 {
        return paths
    }
    index := 1
    helper(n,k,&paths,path,index)
    return paths
}

func helper(n int,k int,paths *[][]int,path []int,index int) {
   if index > n + 1 {
       return
   }
    // terminator
    if len(path) == k{
        tmp := make([]int, k)
        copy(tmp, path)
        *paths = append(*paths, tmp)
        return
    }
    // process current logic
    
    // 不选
    helper(n,k,paths,path,index + 1)
    // 选
    helper(n,k,paths,append(path,index),index + 1)
}

3. 看题解：
func combine(n int, k int) [][]int {
    result := make([][]int, 0)
    helper(1, n, k, []int{}, &result)
    return result
}

func helper(pointer, n, k int, current []int,  result *[][]int) {
    if len(current) == k {
        *result = append(*result, append([]int{}, current...))
        return
    }
    for i := pointer; i <= n; i++ {
        helper(i+1, n, k, append(current, i), result)
    }
}

4. 复杂度分析：
时间复杂度：O(n*n^k)
空间复杂度：O(n)

5. 总结：
5.1: 哈哈哈，有点累，可能是自己能力需要提高了，空余时间尽可能的提高自己的能力哈
