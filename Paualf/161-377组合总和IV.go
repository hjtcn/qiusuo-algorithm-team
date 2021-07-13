给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。
示例 1：
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
示例 2：
输入：nums = [9], target = 3
输出：0
 
提示：
1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000
 
进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？

1. Clearfication:
target = 4,背包为4，顺序不同的被看作是不同的组合，这个地方要注意下

       动规五部曲：
        1. dp[n] : 凑成target 的组合数 dp[target]

        2. dp 公式
           dp[target]：考虑所有的 nums 的组合总和，就是所有的 dp[target - nums[i]] 相加
           所以递推公式：dp[target] += dp[target - nums[i]]

        3. dp 数组如何初始化
            dp[0] = 1
            从dp[i] 的含义上来讲就是：凑成总金额为0的组合数为1
        4. 确定  遍历顺序，因为顺序不同被视作是不同的组合，所以我觉得在这里遍历顺序不那么重要

        5. 举例推导dp数组
        for _,num := range nums {
            for i := num;i <= target;i++ {
                dp[i] = dp[i] + dp[i - num]
            }
        }

        return dp[target]
coding:

func combinationSum4(nums []int, target int) int {
    dp := make([]int,target + 1)
    dp[0] = 1
    for _,num := range nums {
        for i := num;i <= target;i++ {
            dp[i] = dp[i] + dp[i - num]
        }
    }
    return dp[target]
}

上面的执行结果是比较少，应该是把重复的序列排除了，我们换一种遍历方式
func combinationSum4(nums []int, target int) int {
    dp := make([]int,target + 1)
    dp[0] = 1
    for i:=1;i <=target;i++  {
        for _,num := range nums {
            if i - num >= 0 {
                dp[i] = dp[i] + dp[i - num]
            }
        }
    }
    return dp[target]
}

Thinking:
如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
如果允许负数存在的话：
dp初始化的时候，要允许负数的下标存在，判断条件 i - num >= 0 这个条件也应该去掉了

2. 看题解：
看到题解中，提到上述做法是否考虑到选取元素的顺序？
因为外层循环是遍历从1到target的值，内层循环是遍历数组nums的值，在计算dp[i]的值时，nums中的每个小于等于i的元素都可能作为元素之和等于i的排列的最后一个元素。例如，1和3都在nums中，计算dp[4] 的时候，排列的最后一个元素可以是1也可以是3，因此dp[1]和dp[3]都会被考虑到，即不同的顺序都会被考虑到

同理：如果先遍历 nums 1 2 3 的时候
num 从 1 -> 4: 会计算dp[1] + dp[1- 1] / dp[2-1] = dp[1] / dp[2]/ dp[3]
num 从 2 -> 4 会计算 dp[2] + dp[0] / dp[1] + dp[2]
num 从 3 -> 4 会计算 dp[3]  + dp[0]  / dp[1]
看着 dp[1] + dp[3] 出现了2次，但是一开始的dp[1] + dp[3] ,这个时候dp[3] 是0，所以没有起到任何效果
后面 的 dp[3] + dp[1] 才是有值的，所以过滤的重复的组合元素

3. 复杂度分析：
时间复杂度：O(target * len(num))
空间复杂度：O(target)

4. 总结：
4.1: 在数学中组合的实际问题是 1 * x + 2 * y + 3 * z = target,求 x,y,z 的值可能出现的值，然后使用计算机去算的话，需要抽象成计算机可以理解的语言和形式，这里使用动态规划，本质上存储的数据结构是数组，记录计算过的元素
现实中的具体问题 -> 抽象 -> 计算机可以理解的代码/指令

4.2: 都是 for,怎么 for 还是一门学问呀
