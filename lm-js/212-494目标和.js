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
-1000 <= target <= 1000


*/


/*
    思路：我写出来的方法我都有点害怕，哈哈，我用了递归。执行用时:2644ms
         对于+，-来说，要么加，要么减
         因此dfs(i+1,sum+nums[i])
            dfs(i+1,sum-nums[i])
    题解：01背包。
        例如案例
        输入：nums = [1,1,1,1,1], target = 3
        输出：5
        将x设为+号的元素和，sum为元素和，sum-x为-号的元素和
        可以推出 x-(sum-x)=target
                    x=(target+sum)/2
        dp[j]表示填满容量为j的背包，有dp[j]种方法   
        因此我们的终极目标为寻找dp[x]
        
*/

var findTargetSumWays = function(nums, target) {
    let count=0,len=nums.length
    let dfs=(i,sum)=>{
        if(sum==target&&i==len){
            count++;
            return ;
        }
        if(i>=len){
            return;
        }
        dfs(i+1,sum+nums[i])
        dfs(i+1,sum-nums[i])
    }
   dfs(0,0)
   return count
};
/*
    时间复杂度：O(2^N)
    空间复杂度：O(N)
*/


var findTargetSumWays = function(nums, target) {
    let len=nums.length,sum=0
    for(let i=0;i<len;i++){
        sum+=nums[i]
    }
    if(sum<target||(sum+target)%2) return 0
    let rangeX=parseInt((sum+target)/2)
    let dp=Array(rangeX+1).fill(0)
    dp[0]=1
    for(let i=0;i<len;i++){
        for(let j=rangeX;j>=nums[i];j--){
            dp[j]+=dp[j-nums[i]]
        }
    }
    return dp[rangeX]
}

/*
    时间复杂度：O(n*sum)
    空间复杂度：O(sum)
*/

/*
    1.01背包中0和1要区分开去找规律
    2.递归记忆化要找到存储什么？像这道题对于记忆化我就没有多少思路
*/
