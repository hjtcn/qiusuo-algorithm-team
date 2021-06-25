/*
152. 乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。



示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/

/*
    思路：这个和最大和数组就有些不同了。
         需要区分正负数,记录最大值最小值。
         1.当前数为正数，正数*上一步的最大值，会越来越大；正数*上一步的最小值，会越来越小
         2.当前值为负数，负数*上一步的最大值，会越来越小;负数*上一步的最小值，会越来越大。

         动规：
         dp[i][0]代表从0到i的最大值
         dp[i][1]代表从0到i的最小值
        for(let i=0;i<len;i++)
        {   
            if(nums[i]>=0){
                dp[i][0]=max(nums[i],dp[i-1][0]*nums[i])
                dp[i][1]=min(nums[i],dp[i-1][1]*nums[i])
            }
            else{
                dp[i][0]=max(nums[i],dp[i-1][1]*nums[i])
                dp[i][1]=min(nums[i],dp[i-1][0]*nums[i])
            }
            max=max(max,dp[i][0])
        }

    优化：降维
        for(let i=0;i<len;i++){
            let oldMax=prevMax,oldMin=prevMin
            if(nus[i]>=0){
                prevMax=Math.max(nums[i],oldMax*nums[i])
                prevMin=Math.min(nums[i],oldMin*nums[i])
            }
            else{
                  prevMax=Math.max(nums[i],oldMIn*nums[i])
                prevMin=Math.min(nums[i],oldMax*nums[i])
            }
        }

    当然正负判断，也何以合并进行三个数的比较
         prevMax=Math.max(nums[i],oldMax*nums[i],oldMIn*nums[i])
         prevMin=Math.min(nums[i],oldMin*nums[i],oldMax*nums[i])

*/


var maxProduct = function(nums) {
    let len=nums.length,max=nums[0]
    let dp=[]
    for(let i=0;i<len;i++){
        dp[i]=[0,0]
    }
    dp[0][0]=nums[0]
    dp[0][1]=nums[0]
    for(let i=1;i<len;i++){
        if(nums[i]<0){
          dp[i][0]=Math.max((dp[i-1][1]*nums[i]),nums[i])
          dp[i][1]=Math.min(dp[i-1][0]*nums[i],nums[i])
        }
        else{
            dp[i][0]=Math.max(dp[i-1][0]*nums[i],nums[i])
            dp[i][1]=Math.min(dp[i-1][1]*nums[i],nums[i])
        }
        max=Math.max(max,dp[i][0])
       
    } 
    return max
};
/*
    时间复杂度:O(N)
    空间复杂度:O(N)
*/

var maxProduct = function(nums) {
    let len=nums.length,max=nums[0]
    let prevMax=nums[0],prevMin=nums[0]
    for(let i=1;i<len;i++){
        let oldMax=prevMax,oldMin=prevMin
        //合并正负判断
        prevMax=Math.max(nums[i],oldMax*nums[i],oldMin*nums[i])
        prevMin=Math.min(nums[i],oldMin*nums[i],oldMax*nums[i])
        // if(nums[i]>=0){
        //     prevMax=Math.max(nums[i],oldMax*nums[i])
        //     prevMin=Math.min(nums[i],oldMin*nums[i])
        // }
        // else{
        //     prevMax=Math.max(nums[i],prevMin*nums[i])
        //     prevMin=Math.min(nums[i],oldMax*nums[i])
        // }
        max=Math.max(prevMax,max)
    } 
    return max
};

/*
    时间复杂度:O(N)
    空间复杂度:O(1)
*/

/*
    思考：今天的正负判断，有点吓着了。
        当时考虑的是从i到j，不断更新max
        没有想到直接从0到i给记录住。
*/