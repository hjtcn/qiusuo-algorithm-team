/*
416. 分割等和子集
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

 

示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
 

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/
/*
    思路：01背包。确认背包容量为sum/2
        1.求背包容量为sum/2的情况下，能最多装多少重量(此题即最多数字)？
        （1）确认dp意义
         dp[j]代表背包容量为j时，最多能装多少数字，价值为数字大小。
        （2）确认状态转移方程.第一个nums[i]为占位，第二个为价值
         dp[j]=Max(dp[j],dp[j-nums[i]]+nums[i])
        （3）dp初始化
         dp=Array(sum/2+1).fill(0)
        （4）遍历顺序
         还是按照习惯，先数字，再背包容量
         for(let i=0;i<len;i++){
             for(let j=target;j>=nums[i];j--){
                 dp[j]=Max(dp[j],dp[j-nums[i]]+nums[i])
             }
         }
        （5）举例推导
        2.突然想起来dp也可以记录是否能构建出sum/2的和
         (1).dp[j]为可构建出和为j的次数
         (2).dp[j]=dp[j]+dp[j-nums[i]]
         (3).dp=Array(target+1).fill(0)
             dp[0]=1
         (4).遍历顺序，先数字数组遍历，再背包容量遍历
         (5).举例说明推导
         这种方式，dp的意义不同了，利用递推的方式。和第一种明显构建01背包的感觉还是不太一样的。

        3.递归
        (1).边界是什么？
         if(index>=array.length){
             return false
         }
        (2).目标返回是什么？
         if(sum==target){
             return true
         }
        (3).题目思路是什么？
         当前数可选 dfs(sum+nums[i],i++)
              可不选 dfs(sum,i++)

        记忆化
        (1)记忆什么？
        记忆索引所对应的返回值。
        (2)返回值应为什么？
        无论当前值怎么操作，只有抵达target，才能返回true，其他都为false
        return dfs(sum+nums[i],i++)||dfs(sum,i++)
        (3)怎么记录返回值?
        map.set(i,dfs(sum+nums[i],i++)||dfs(sum,i++))
        上面map的key值更有唯一性一些，如:key=sum+'&'+i
        
        虽然理解了大致思路，但还是遇到了两个细节坑点，导致了超时：
        (1)边界还是不够准确。如果sum>target，也可以提前返回
       （2)map.get(key)的耗时,会比map.has(key)长


    



*/

var canPartition = function (nums) {
    let sum=0,len=nums.length
    for(let i=0;i<len;i++){
        sum+=nums[i]
    }
    if(sum%2) return false
    let target=parseInt(sum/2)
    let dp=Array(target+1).fill(0)
    for(let i=0;i<len;i++){
        for(let j=target;j>=nums[i];j--){
            dp[j]=Math.max(dp[j],dp[j-nums[i]]+nums[i])
        }
    }
    if(dp[target]==target){
        return true
    }
    else{
        return false
    }
  };


  var canPartition = function (nums) {
    let sum=0,len=nums.length
    for(let i=0;i<len;i++){
        sum+=nums[i]
    }
    if(sum%2) return false
    let target=parseInt(sum/2)
    let dp=Array(target+1).fill(0)
    dp[0]=1
    for(let i=0;i<len;i++){
        for(let j=target;j>=nums[i];j--){
            dp[j]=dp[j]+dp[j-nums[i]]
        }
    }
    return dp[target]
  };


var canPartition = function (nums) {
    let sum=0,len=nums.length
    for(let i=0;i<len;i++){
        sum+=nums[i]
    }
    if(sum%2) return false
    let map=new Map()
    let target=parseInt(sum/2)
    let dfs=(sum,index)=>{
        if(index==len||sum>target){
            return false
        }
        if(sum==target){
            return true
        }
        let key=sum+'&'+index
        if(map.has(key)){
            return map.get(key)
        }
        let res=dfs(sum+nums[index],index+1)||dfs(sum,index+1)
        map.set(key,res)
        return res
    }
    return dfs(0,0)
  };

  /*
    时间复杂度：O(N*target)
    空间复杂度：O(target)
  */

  /*
    思考：今天的题比前两天要更基础一些，也更方便我们去总结理解。
  */
