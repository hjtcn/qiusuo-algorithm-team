给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
示例 1:
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数； 
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。

没有思路，自己也没想着怎么动手，还是有点害怕了，其实最简单暴力法是也可以做出来的，自己都没有尝试，诶。。。

暴力法：
将数组扩大一倍，然后查找比它大的元素，接一半然后去查找

func nextGreaterElements(nums []int) []int {
    ret := make([]int,len(nums))
    doubleNum := make([]int, 2 * len(nums))    
    // 扩大两倍
    round := 0
    index := 0
    for round < 2 {
        round++
         for i := 0;i < len(nums);i++ {
             doubleNum[index] = nums[i]
             index++
         }
    }
   //fmt.Println(doubleNum)
    for i := 0;i < len(nums);i++ {
        ret[i] = -1
        for j := i + 1;j < len(doubleNum);j++ {
            if doubleNum[j] > nums[i] {
                ret[i] = doubleNum[j]
                break
            }
        }
    }
    return ret
}

复杂度分析：
时间复杂度：O(n) ^2 ：两次for 循环
空间复杂度：O(n) : 开辟了一个新的数组，将数据放入到数组中

使用单调栈遍历两边数组：

func nextGreaterElements(nums []int) []int {
    if len(nums) == 0 {
        return nil
    }
    res := make([]int, len(nums))
    for i := 0;i < len(res);i++ {
        res[i] = -1
    }
    stack := []int{0}
    i := 0
    round := 0
    for round < 2 {
        i++
        if i == len(nums) {
            i = 0
            round++
        }
        for len(stack) != 0 && nums[i] > nums[stack[len(stack) - 1]] {
            res[stack[len(stack) - 1]] = nums[i]
            stack = stack[:len(stack) - 1]
        }
        if res[i] == -1 {
            stack = append(stack, i)
        }
    }
    
    return res
}

复杂度分析：
时间复杂度：O(n) ：遍历两边数组，时间复杂度仍是O(n)
空间复杂度：O(n) : 使用栈

总结：
1. 官方题解没有看懂，里面这个nums[(i + j) % nums.length]，没懂为什么这样写，并且中文的官方题解只有1种，英文的有3种。写的代码也是云里雾里的，感觉还没有其他人的代码好理解，也不知道是不是错觉了，或许有的时候官方题解真的不咋样。

2. 一道题如果没有好的思路，就想从暴力的办法开始写，写出暴力的办法以后再去想怎么优化，不要害怕，动手分析
