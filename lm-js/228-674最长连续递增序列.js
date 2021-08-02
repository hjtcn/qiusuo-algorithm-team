/*
674. 最长连续递增序列
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

 

示例 1：

输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。 
示例 2：

输入：nums = [2,2,2,2,2]
输出：1
解释：最长连续递增序列是 [2], 长度为1。
 

提示：

1 <= nums.length <= 104
-109 <= nums[i] <= 109

*/

/*
    思路：这种连续递增的，还是更倾向于贪心吧，时间复杂度为O(1)
        就是遍历，更新最大连续递增长度。一旦不递增，则从长度赋值为1，接着更新

        动态规划，思路基本一致，区别仅在于
        1.dp数组的意义：
          dp[i]0～i的最长递增子序列长度
        2.递推公式
         dp[i]=dp[i-1]+1(当然递增的时候)
        3.初始化
        dp数组初始化为1
        dp=Array(len).fill(1)
        4.遍历顺序，只有一层，按习惯从前往后吧
        5.举例
*/

var findLengthOfLCIS = function(nums) {
    let len=nums.length
    let count=1,max=1
    for(let i=1;i<len;i++){
        if(nums[i]>nums[i-1]){
            count++;
        }
        else{
            count=1
        }
        max=Math.max(count,max)
    }
    return max
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

var findLengthOfLCIS = function(nums) {
    let len=nums.length
    let dp=Array(len).fill(1),max=1
    for(let i=1;i<len;i++){
        if(nums[i]>nums[i-1]){
            dp[i]=dp[i-1]+1
        }
        max=Math.max(dp[i],max)
    }
    return max
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    简单题，调整调整心情，接下来会有些忙呀。
*/