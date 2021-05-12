给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
示例 1:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:
输入: [2,3,0,1,4]
输出: 2
 
提示:
1 <= nums.length <= 1000
0 <= nums[i] <= 105

1. Clearfication:
1.1: 是不是可以跳跃到最后一个位置
1.2: 到达数组最后一个位置最少的跳跃数
     最少的跳跃数这个怎么去定义它呢？
2 2 3 1 4
分析了这个case，2 -> 2 -> 1 -> 4 3步
2->3->4 2步，怎么去用代码去模式这个过程呢？ =》 没有想好

2. Coding:

3. 看题解：
蛮巧妙的思路：
多个位置通过跳跃能够到达最后一个位置，如何进行选择呢？
可以 [贪心] 地选择距离最后一个位置最远的位置，也就是对应下标最小的位置。
从左到右遍历数组，选择第一个满足要求的位置。
找到最后一步跳跃前所在的位置之后，我们继续贪心地寻找倒数第二步跳跃所在的位置，直到找到数组的开始位置。
func jump(nums []int) int {
    position := len(nums) - 1
    steps := 0
    
    for position > 0 {
        for i := 0;i < position;i++ {
            if i + nums[i] >= position {
                position = i
                steps++
                break
            }
        }
    }
    
    return steps
}

也联想到了这种bad case：
【2，0，0，3，1，4】 这种是 不合理的输入情况

从前往后跳：
比较巧妙的地方在
 i == end {
end = maxPosition
steps++
}
func jump(nums []int)int {
    length := len(nums)
    end := 0
    maxPosition := 0
    steps := 0
    
    for i := 0;i < length - 1;i++ {
        maxPosition = max(maxPosition, i + nums[i])
        if i == end {
            end = maxPosition
            steps++
        }
    }
    
    return steps
}

func max(x,y int) int {
    if x > y {
        return x
    }
    
    return y
}

看了题解，好像第一种解法确实有bfs的影子，还看到使用memo记忆化数组的方法和动态规划的方式，发现自己现在好像并没有真正理解这道题的意思。。。

4. 复杂度分析：
时间复杂度：第一种方法：O(n*n), 第二中方法：O(n)
空间复杂度：O(1)

5. 总结：
5.1: 跳一跳还是蛮巧妙的，更新边界值去处理还是挺厉害的
5.2: 算法真的是灵魂，同样是一维数组，你写的代码和别人写的代码不是一个代码
