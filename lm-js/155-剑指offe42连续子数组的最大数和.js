// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

// 要求时间复杂度为O(n)。

//  

// 示例1:

// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//  

// 提示：

// 1 <= arr.length <= 10^5
// -100 <= arr[i] <= 100
// 注意：本题与主站 53 题相同：https://leetcode-cn.com/problems/maximum-subarray/


/*
    思路：动态规划没思路。画了半天状态转移，找不出来以那个值，后来看了看之前的题解，发现它经历了两次Math.max()的更新。虽然代码很少，但是自己去构建思路还是不容易的
    dp(x)=max(dp[x-1]+nums[i],nums[i]) //去判断是和值作为最大值还是当前值作为最大值。
    最后max=max(dp(x),max)...不断记录最大值，只用于返回，和状态转移是无关的。。。我思考的时候总是把dp,max弄混，最后思路混乱了。
        这次知道是完全没考虑之前是咋做的，就想画状态转移方程，但是没有思路，这时候才明白，我们做的确实是简单题呀，扩展类的基本还没怎么接触。好吧，动规没入门，继续训练

                           
*/
var maxSubArray = function(nums) {
    let len=nums.length,dp=0,max=-Infinity;
    for(let i=0;i<len;i++){
        dp=Math.max(dp+nums[i],nums[i]) 
        max=Math.max(dp,max)
    }
    return max
};

//贪心。保证和值最大的形式去进行遍历
var maxSubArray = function(nums) {
    let len=nums.length,max=-Infinity,sum=0
    for(let i=0;i<len;i++){
        if(sum>0){
            sum+=nums[i]
        }
        else{
            sum=nums[i]
        }
        //不断更新目标最大值
        max=Math.max(sum,max)
    }
    return max
}

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/