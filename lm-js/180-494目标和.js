/*
494. 目标和
给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

 

示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1
 

提示：

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 100

*/


/*
    思路：尝试枚举了，但是木有成功.
    题解：看了枚举的思路，联想到了二叉树中和为某一值的路径这种类型题。觉得没有做出来还是比较可惜的。

        01背包简直要崩溃，模模糊糊把思路走通了。。。。
        刚开始是数学题，要将+，-对比为01背包
        设x为前置为+的元素和，sum-x为前置为-的元素和
        x-(sum-x)=target
        x=(target+sum)/2
        此时问题就变为在nums中找出和为x的可能性
        for(let i=0;i<nums.length;i++){
            for(let j=x;j>=nums[i];j--){
                dp[j]+=dp[j-nums[i]]
            }
        }
*/


var findTargetSumWays = function(nums, target) {
    let count=0
    let dfs=(nums,i,sum,target)=>{
        if(i==nums.length){
            if(sum==target){
                count++;
            }
        }
        else{
            dfs(nums,i+1,sum+nums[i],target)
            dfs(nums,i+1,sum-nums[i],target)
        }
    }
    dfs(nums,0,0,target)
    return count
};

/*
    时间复杂度：O(2^n)
    空间复杂度：O(N)
*/



var findTargetSumWays = function(nums, target) {
    let len=nums.length,sum=0
    for(let i=0;i<len;i++){
        sum+=nums[i]
    }
    if(sum<target||(sum+target)%2) return 0
    let makeCount=(sum+target)/2 
    let dp=Array(makeCount+1).fill(0)
    dp[0]=1
   for(let i=0;i<len;i++){
       for(let j=makeCount;j>=nums[i];j--){      
           dp[j]=dp[j]+dp[j-nums[i]]
       }
   }   
    return dp[makeCount]
};

/*
    时间复杂度：O(n*sum)
    空间复杂度：O(sum)
*/

/*
    01背包模块化，思考并转化为01背包模型
*/