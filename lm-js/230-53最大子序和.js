/*
53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

 

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [0]
输出：0
示例 4：

输入：nums = [-1]
输出：-1
示例 5：

输入：nums = [-100000]
输出：-100000
 

提示：

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105
 

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/

/*
    贪心：不从元素个数出发，从和出发。
        如果和小于0，它就不配加nums[i],因为越加越小嘛。
        因此sum<0,则sum=nums[i]
        sum>=0.则sum+=nums[i]
    动规。dp也是从和出发
    dp=Math.max(dp+nums[i],nums[i])
    到底是(dp+当前值)大，还是当前值本身就比较大一些，这样的更新过程。
    同时更新最大值max,跳出循环返回即可

*/

var maxSubArray = function(nums) {
    let len=nums.length
    let sum=0,max=-Infinity
    for(let i=0;i<len;i++){
        if(sum>0){
            sum+=nums[i]
        }
        else{
            sum=nums[i]
        }
        max=Math.max(sum,max)
    }
    return max
};

var maxSubArray = function(nums) {
    let len=nums.length
    let max=-Infinity
    let dp=0
    for(let i=0;i<len;i++){
       dp=Math.max(dp+nums[i],nums[i])
       max=Math.max(dp,max)
    }
    return max
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)

*/

/*
    这道题并不是那么简单，和我平常的思路是反向的，一般我思考和，是要不要加当前值，而这道题是要思考要不要加之前的和。
*/