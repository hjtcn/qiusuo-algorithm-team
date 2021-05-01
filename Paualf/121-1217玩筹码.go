数轴上放置了一些筹码，每个筹码的位置存在数组 chips 当中。
你可以对 任何筹码 执行下面两种操作之一（不限操作次数，0 次也可以）：
• 将第 i 个筹码向左或者右移动 2 个单位，代价为 0。
• 将第 i 个筹码向左或者右移动 1 个单位，代价为 1。
最开始的时候，同一位置上也可能放着两个或者更多的筹码。
返回将所有筹码移动到同一位置（任意位置）上所需要的最小代价。
 
示例 1：
输入：chips = [1,2,3]
输出：1
解释：第二个筹码移动到位置三的代价是 1，第一个筹码移动到位置三的代价是 0，总代价为 1。
示例 2：
输入：chips = [2,2,2,3,3]
输出：2
解释：第四和第五个筹码移动到位置二的代价都是 1，所以最小总代价为 2。
提示：
• 1 <= chips.length <= 100
• 1 <= chips[i] <= 10^9

1. Clearfication:
这道题目，题目没有看懂，第四和第五个筹码到位置二的代价都是1 ？ 没看懂。。。

2. Coding:

3. 看题解：
func minCostToMoveChips(position []int) int {
    odd,even := 0,0
    
    for i := 0;i < len(position);i++ {
        if position[i] % 2 == 0 {
            even++
        }else if (position[i] % 2) != 0 {
            odd++
        }
    }
    
    return min(even,odd)
}

func min(x,y int) int {
    if x < y {
        return x
    }
    return y
}

https://leetcode-cn.com/problems/minimum-cost-to-move-chips-to-the-same-position/solution/xian-li-jie-ti-yi-zai-li-jie-dai-ma-si-lu-by-athen/
看了这个题解明白了，[2,2,2,3,3] 这个为什么移动是2了
第1个筹码放第2个位置，第2个筹码放第2个位置，第3个筹码放第2个位置，第4个筹码放第3个位置，第5个筹码放第3个位置，那么这就表示，第2个位置上有3个筹码，第3个位置上有2个筹码，其他位置上没有筹码
思路：移动2个位置不需要代价，那么奇数位置到奇数位置不用代价，偶数位置到偶数位置不用代价，那就统计奇数位置和偶数位置的个数，相当于把所有奇数位置放一起，所有偶数放一起，然后比较奇数的少还是偶数的少，将少的个数移动多的个数位置上就可以了
尽量使用for range 

func minCostToMoveChips(chips []int)int {
    odd,even := 0,0
    
    for _,v := range chips {
        if v % 2 == 0 {
            even++
        }else {
            odd++
        }
    }
    
    return min(odd,even)
}

func min(x,y int) int {
    if x < y {
        return x
    }
    return y
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

5. 总结：
5.1: 题目没有理解清楚，理解题目是把题目做对的很关键的一步，项目中也是，把需求分析清楚是把需求做好的最关键的一部分
