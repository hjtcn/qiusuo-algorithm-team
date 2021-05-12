/*
45. 跳跃游戏 II
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

*/

/*
    思路：哈哈，今天没有思路，能写出来昨天能跳到的最远距离，但是边界和跳跃次数更新实在没思路。看题解吧。
    
    题解：
    方法1:接着说能跳到最远距离，边界时控制如果i==前一次的最远距离值。说明要开始跳下一次了，此时更新最远距离值，并跳跃次数加1.

    方法2:倒序并且不断更新目标位置。如果当前下标i+val能大于等于目标位置，则跳跃步数+1，目标位置也可以设置为i

    方法3:动规，j代表起始位置，i代表结束位置。
    两层for循环，j<i,如果当前值nums[j]可以达到跳跃长度，即nums[j]>j-i,则更新跳跃次数最小值最小值dp[i]=min(dp[i],dp[j]+1).
*/

var jump = function (nums) {
    let max = 0, end = 0, count = 0
    for (let i = 0; i < nums.length - 1; i++) {
        max = Math.max(max, i + nums[i])
        if (i == end) {
            // 步子一旦跨的比较大，就贪心不起来了。哈哈哈
            end = max
            count++
        }
    }
    return count
}
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/


var jump = function (nums) {
    let target = nums.length - 1,count=0
    while (target > 0) {
        for (let i = 0; i < target; i++) {
            if (i + nums[i] >= target) {
                count++;
                target = i;
                break;
            }
        }
    }
    return count
}

/*
    时间复杂度：O(N*2)
    空间复杂度：O(1)
*/
//动规

var jump = function (nums) {
    let len=nums.length;
    let dp=[]
    dp[0]=0
    for(let i=1;i<len;i++){
        dp[i]=Infinity
        for(let j=0;j<i;j++){
            if(nums[j]>=i-j)
            {
                dp[i]=Math.min(dp[i],dp[j]+1)
            }
        }
    }
    return dp[len-1]
}

/*
    时间复杂度：O(N*2)
    空间复杂度：O(N)
*/

/*
    看到这个题，真的觉得可适合做动规了，但是我脑子里对于两层遍历还是一层遍历来回犹豫，模型也没画出来，
    好吧，万物皆可动规，不要陷在一种题型里面，但做每道题的时候都尽力构建一种题型的模型.
*/

